package debug

import "testing"

// 单一常驻会话:目标身份判定/空闲宿主复用/环境克隆的纯逻辑单测
// (SSH 登录等依赖真实 T100,不在此覆盖)

func cfgFor(host, zone string) *Config {
	c := &Config{}
	c.SSH = SSHConfig{Host: host, Port: 22, User: "u"}
	c.Zone = zone
	c.fillDefaults()
	return c
}

func fakeIdleSession(id, host, zone string) *Session {
	return &Session{
		ID:       id,
		cfg:      cfgFor(host, zone),
		state:    StateIdle,
		Module:   "oldm",
		Prog:     "oldp",
		bps:      map[int]*Breakpoint{},
		srcCache: map[string]srcCacheEntry{},
	}
}

func TestSessionSameTarget(t *testing.T) {
	s := fakeIdleSession("s1", "h1", "36")
	if !s.sameTarget(cfgFor("h1", "36")) {
		t.Fatal("同 host/zone 应判定为同一目标")
	}
	if s.sameTarget(cfgFor("h2", "36")) {
		t.Fatal("不同 host 不应判定为同一目标")
	}
	if s.sameTarget(cfgFor("h1", "35")) {
		t.Fatal("不同 zone 不应判定为同一目标")
	}
	if s.sameTarget(nil) {
		t.Fatal("nil 配置不应判定为同一目标")
	}
}

func TestPrepareSessionReusesIdleHost(t *testing.T) {
	m := NewManager(cfgFor("h1", "36"))
	live := fakeIdleSession("s1", "h1", "36")
	m.mu.Lock()
	m.sessions["s1"] = live
	m.mu.Unlock()

	sess, err := m.prepareSession(cfgFor("h1", "36"), "m1", "p1", "", "", "", "")
	if err != nil {
		t.Fatalf("prepareSession: %v", err)
	}
	if sess != live {
		t.Fatal("空闲同目标会话应被复用,而不是新建")
	}
	if sess.Module != "m1" || sess.Prog != "p1" {
		t.Fatalf("复用时应刷新本轮运行参数: module=%q prog=%q", sess.Module, sess.Prog)
	}
	if m.Get("s1") == nil {
		t.Fatal("复用后会话应仍登记在管理器")
	}
}

func TestPrepareSessionRejectsActiveRun(t *testing.T) {
	m := NewManager(cfgFor("h1", "36"))
	active := fakeIdleSession("s1", "h1", "36")
	active.mu.Lock()
	active.state = StateStopped
	active.mu.Unlock()
	m.mu.Lock()
	m.sessions["s1"] = active
	m.mu.Unlock()

	if _, err := m.prepareSession(cfgFor("h1", "36"), "m1", "p1", "", "", "", ""); err == nil {
		t.Fatal("同目标活跃会话(stopped)应报错,要求先结束当前调试")
	}
}

func TestTopentForRun(t *testing.T) {
	c := cfgFor("h1", "36")
	c.DB = &DBConfig{Ent: 7}
	s := fakeIdleSession("s1", "h1", "36")
	s.cfg = c
	if got := s.topentForRun(); got != "7" {
		t.Fatalf("无手动设置时应回退配置 DB.Ent,got %q", got)
	}
	s.mu.Lock()
	s.topentOverride = "99"
	s.mu.Unlock()
	if got := s.topentForRun(); got != "99" {
		t.Fatalf("手动设置应优先于配置企业,got %q", got)
	}
	s2 := fakeIdleSession("s2", "h1", "36")
	if got := s2.topentForRun(); got != "" {
		t.Fatalf("无 DB 且无手动设置应返回空(沿用登录默认),got %q", got)
	}
	if got := s.TopentOverride(); got != "99" {
		t.Fatalf("TopentOverride 应返回手动值,got %q", got)
	}
}

func TestCloneEnvAndEnvName(t *testing.T) {
	c := &Config{
		SSH: SSHConfig{Host: "top", Port: 22, User: "u"},
		Envs: []NamedEnv{
			{Name: "E1", SSHConfig: SSHConfig{Host: "e1h", Port: 22, User: "u1"}, Zone: "35", DB: &DBConfig{Ent: 7}},
			{Name: "E2", SSHConfig: SSHConfig{Host: "e2h", Port: 22, User: "u2"}, Zone: "36"},
		},
	}
	c.fillDefaults()

	clone := c.CloneEnv("E1")
	if clone == nil {
		t.Fatal("CloneEnv 应命中 E1")
	}
	if clone.SSH.Host != "e1h" || clone.Zone != "35" || clone.ActiveEnv != "E1" {
		t.Fatalf("CloneEnv 字段不符: %+v", clone.SSH)
	}
	if clone.DB == nil || clone.DB.Ent != 7 {
		t.Fatal("CloneEnv 应带环境专属 DB 配置")
	}
	if c.SSH.Host != "top" {
		t.Fatal("CloneEnv 不应改动原配置")
	}
	if c.CloneEnv("Nope") != nil {
		t.Fatal("未命中环境应返回 nil")
	}
	if got := cfgFor("x", "36").EnvName(); got != "x-36" {
		t.Fatalf("无 activeEnv 时环境名应为 host-zone,got %q", got)
	}
}
