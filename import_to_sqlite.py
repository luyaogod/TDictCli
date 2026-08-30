#!/usr/bin/env python3
"""
ERP 数据字典导入 SQLite 脚本
------------------------------
从 data/ 目录读取 dzea_t, dzeal_t, dzeb_t, dzebl_t, dzed_t 五个 CSV 文件，
1. 先查询展示五张表自身的元数据（自引用记录）
2. 在 SQLite 中创建同名表并完整导入所有数据
3. 根据 dzed_t 中的 PK 定义创建索引
"""

import csv
import sqlite3
import os
import sys
import time

# ============================================================
# 配置
# ============================================================
DATA_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)), "data")
DB_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)), "erp_data.db")

TABLES = ["dzea_t", "dzeal_t", "dzeb_t", "dzebl_t", "dzed_t"]

# 每张表的 CSV 文件编码（通过抽样检测确定）
# 这些文件使用系统默认编码，在 Windows 上通常是 cp950 (BIG5) 或 utf-8
FILE_ENCODING = "utf-8"

# 批量插入的行数
BATCH_SIZE = 10000


def detect_encoding(filepath):
    """检测文件编码"""
    # 先尝试 utf-8
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            f.read(1024)
        return "utf-8"
    except UnicodeDecodeError:
        pass
    # 尝试 utf-8-sig (BOM)
    try:
        with open(filepath, "r", encoding="utf-8-sig") as f:
            f.read(1024)
        return "utf-8-sig"
    except UnicodeDecodeError:
        pass
    # 尝试 cp950 (BIG5 繁体中文)
    try:
        with open(filepath, "r", encoding="cp950") as f:
            f.read(1024)
        return "cp950"
    except UnicodeDecodeError:
        pass
    # 尝试 gbk (简体中文)
    try:
        with open(filepath, "r", encoding="gbk") as f:
            f.read(1024)
        return "gbk"
    except UnicodeDecodeError:
        pass
    # 回退到 latin-1（不会失败）
    return "latin-1"


def read_csv(filepath):
    """
    读取逗号分隔的 CSV 文件，返回 (headers, rows) 元组。
    文件特点：
    - 分隔符为逗号 (,)
    - 部分字段用双引号包裹
    - 部分字段值末尾含有 tab 字符 (\\t)，需清理
    - 行尾为 \\r\\n
    """
    encoding = detect_encoding(filepath)
    print(f"  读取 {os.path.basename(filepath)} (编码: {encoding})...")

    with open(filepath, "r", encoding=encoding, newline="") as f:
        # 使用逗号分隔，QUOTE_ALL 处理所有带引号字段
        reader = csv.reader(f, delimiter=",", quoting=csv.QUOTE_ALL)
        headers = next(reader)
        # 清理 header：去除引号残留和空白
        headers = [h.strip().strip('"').strip() for h in headers]
        rows = list(reader)
        # 清理每个字段：去除引号、tab字符、首尾空白
        cleaned_rows = []
        for row in rows:
            cleaned_row = []
            for v in row:
                # 去除首尾空白和 tab
                v = v.strip()
                v = v.strip('\t')
                # 去除引号包裹
                if v.startswith('"') and v.endswith('"'):
                    v = v[1:-1]
                cleaned_row.append(v)
            cleaned_rows.append(cleaned_row)
    print(f"    列数: {len(headers)}, 数据行数: {len(cleaned_rows)}")
    return headers, cleaned_rows


def create_table(conn, table_name, headers):
    """创建 SQLite 表，所有列使用 TEXT 类型"""
    # 使用双引号包裹列名以保留原始大小写和特殊字符
    columns = ", ".join(f'"{h}" TEXT' for h in headers)
    sql = f'CREATE TABLE IF NOT EXISTS "{table_name}" ({columns})'
    conn.execute(sql)
    print(f"  创建表 {table_name} ({len(headers)} 列)")


def import_data(conn, table_name, headers, rows):
    """批量导入数据到 SQLite 表"""
    placeholders = ", ".join("?" for _ in headers)
    quoted_cols = ", ".join(f'"{h}"' for h in headers)
    sql = f'INSERT INTO "{table_name}" ({quoted_cols}) VALUES ({placeholders})'

    cursor = conn.cursor()
    total = len(rows)
    for i in range(0, total, BATCH_SIZE):
        batch = rows[i : i + BATCH_SIZE]
        cursor.executemany(sql, batch)
        conn.commit()
        pct = min(100, (i + len(batch)) * 100 // total)
        print(f"\r  导入 {table_name}: {i + len(batch)}/{total} ({pct}%)", end="")
    print()  # 换行


def show_self_metadata(conn):
    """查询并展示五张表自身的元数据记录"""
    print("=" * 80)
    print("Step 1: 五张表自引用元数据查询")
    print("=" * 80)

    # --- dzea_t 中的五表记录 ---
    print("\n>>> dzea_t（資料表主檔）中五表的定义记录：")
    target_names = "', '".join(TABLES)
    rows = conn.execute(
        f"""SELECT dzea001, dzea002, dzea003, dzea004, dzea005, dzea006, dzea007, dzea008
            FROM dzea_t WHERE dzea001 IN ('{target_names}') ORDER BY dzea001"""
    ).fetchall()
    print(f"  {'表名':<12} {'表说明':<24} {'模块':<6} {'类型':<6} {'需翻译':<8} {'系统表':<8} {'含共用字段':<12} {'定义代号':<12}")
    print(f"  {'-'*12} {'-'*24} {'-'*6} {'-'*6} {'-'*8} {'-'*8} {'-'*12} {'-'*12}")
    for r in rows:
        print(f"  {r[0]:<12} {r[1]:<24} {r[2]:<6} {r[3]:<6} {r[4]:<8} {r[5]:<8} {r[6]:<12} {(r[7] or ''):<12}")

    # --- dzed_t 中的五表 PK 定义 ---
    print("\n>>> dzed_t（資料表鍵值檔）中五表的 PK 定义：")
    rows = conn.execute(
        f"""SELECT dzed001, dzed002, dzed003, dzed004, dzed005, dzed006
            FROM dzed_t WHERE dzed001 IN ('{target_names}') ORDER BY dzed001"""
    ).fetchall()
    print(f"  {'表名':<12} {'键名':<16} {'类型':<6} {'键值字段':<50} {'外键表':<12} {'外键字段':<30}")
    print(f"  {'-'*12} {'-'*16} {'-'*6} {'-'*50} {'-'*12} {'-'*30}")
    for r in rows:
        print(f"  {r[0]:<12} {r[1]:<16} {r[2]:<6} {(r[3] or '')[:48]:<50} {(r[4] or ''):<12} {(r[5] or '')[:28]:<30}")

    # --- dzeb_t 中五表各自的字段数量统计 ---
    print("\n>>> dzeb_t（資料表欄位檔）中五表各自的字段数量：")
    for t in TABLES:
        count = conn.execute("SELECT COUNT(*) FROM dzeb_t WHERE dzeb001 = ?", (t,)).fetchone()[0]
        print(f"  {t}: {count} 个字段定义")

    # --- dzeal_t 中五表的多语言描述 ---
    print("\n>>> dzeal_t（資料表多語言檔）中五表的多语言描述：")
    rows = conn.execute(
        f"""SELECT dzeal001, dzeal002, dzeal003 FROM dzeal_t
            WHERE dzeal001 IN ('{target_names}') ORDER BY dzeal001, dzeal002"""
    ).fetchall()
    print(f"  {'表名':<12} {'语言':<8} {'表名(翻译)':<30}")
    print(f"  {'-'*12} {'-'*8} {'-'*30}")
    for r in rows:
        print(f"  {r[0]:<12} {r[1]:<8} {r[2]:<30}")


def show_verification(conn, csv_row_counts):
    """验证导入结果"""
    print("\n" + "=" * 80)
    print("Step 3: 验证导入结果")
    print("=" * 80)

    print(f"\n{'表名':<12} {'CSV行数':>10} {'SQLite行数':>12} {'状态':>8}")
    print(f"{'-'*12} {'-'*10} {'-'*12} {'-'*8}")
    all_ok = True
    for table_name in TABLES:
        csv_count = csv_row_counts[table_name]
        db_count = conn.execute(f'SELECT COUNT(*) FROM "{table_name}"').fetchone()[0]
        status = "[OK]" if csv_count == db_count else "[MISMATCH!]"
        if csv_count != db_count:
            all_ok = False
        print(f"{table_name:<12} {csv_count:>10} {db_count:>12} {status:>8}")

    # 数据库文件大小
    db_size = os.path.getsize(DB_PATH)
    print(f"\n数据库文件: {DB_PATH}")
    print(f"文件大小: {db_size / (1024*1024):.1f} MB")

    # 列出所有表
    print(f"\n数据库中的表:")
    tables = conn.execute(
        "SELECT name FROM sqlite_master WHERE type='table' ORDER BY name"
    ).fetchall()
    for t in tables:
        row_count = conn.execute(f'SELECT COUNT(*) FROM "{t[0]}"').fetchone()[0]
        print(f"  {t[0]}: {row_count} 行")

    if all_ok:
        print("\n[OK] All tables imported successfully!")
    else:
        print("\n[WARN] Some table row counts don't match, please check!")


def main():
    # 处理 Windows 控制台编码问题
    if sys.platform == "win32":
        try:
            sys.stdout.reconfigure(encoding="utf-8")
        except Exception:
            pass

    start_time = time.time()
    print("ERP Data Dict -> SQLite Import Tool")
    print(f"Data dir: {DATA_DIR}")
    print(f"Output DB: {DB_PATH}")
    print()

    # ============================================================
    # 读取所有 CSV 文件
    # ============================================================
    print("Phase 1: 读取 CSV 文件")
    print("-" * 40)
    all_data = {}
    csv_row_counts = {}
    for table_name in TABLES:
        filepath = os.path.join(DATA_DIR, f"{table_name}.csv")
        if not os.path.exists(filepath):
            print(f"  [ERROR] file not found: {filepath}")
            sys.exit(1)
        headers, rows = read_csv(filepath)
        all_data[table_name] = (headers, rows)
        csv_row_counts[table_name] = len(rows)

    # ============================================================
    # 创建 SQLite 数据库并导入数据
    # ============================================================
    print(f"\nPhase 2: 创建 SQLite 数据库并导入数据")
    print("-" * 40)

    # 删除旧数据库
    if os.path.exists(DB_PATH):
        os.remove(DB_PATH)
        print(f"  已删除旧数据库: {DB_PATH}")

    conn = sqlite3.connect(DB_PATH)
    conn.execute("PRAGMA journal_mode=WAL")  # 提升写入性能
    conn.execute("PRAGMA synchronous=OFF")

    for table_name in TABLES:
        headers, rows = all_data[table_name]
        print(f"\n处理 {table_name}...")
        create_table(conn, table_name, headers)
        import_data(conn, table_name, headers, rows)

    # ============================================================
    # 创建索引
    # ============================================================
    print(f"\nPhase 3: 创建索引")
    print("-" * 40)

    # 根据 dzed_t 中五表的 PK 定义创建索引
    dzed_rows = all_data["dzed_t"][1]
    # dzed_t 的 header 列顺序: dzedstus, dzed001, dzed002, dzed003, dzed004, dzed005, dzed006, ...
    # 我们需要 dzed001(表名), dzed002(键名), dzed003(键类型), dzed004(字段列表)
    for row in dzed_rows:
        # row 是按位置索引的，对应 CSV header 的顺序
        # 我们通过 headers 建立映射
        pass

    # 用 SQL 查询方式创建索引，更可靠
    pk_rows = conn.execute(
        """SELECT dzed001, dzed002, dzed004 FROM dzed_t
           WHERE dzed001 IN ('dzea_t','dzeal_t','dzeb_t','dzebl_t','dzed_t')
           AND dzed003 = 'P'"""
    ).fetchall()

    for table_name, key_name, key_fields in pk_rows:
        if not key_fields:
            continue
        # 将逗号分隔的字段列表转为 SQL
        fields = [f.strip().strip('"') for f in key_fields.split(",")]
        quoted_fields = ", ".join(f'"{f}"' for f in fields)
        # 创建唯一索引（SQLite 的 UNIQUE INDEX 等同于 PK 约束）
        idx_sql = f'CREATE UNIQUE INDEX IF NOT EXISTS "{key_name}" ON "{table_name}" ({quoted_fields})'
        try:
            conn.execute(idx_sql)
            print(f"  [OK] create index {key_name} ON {table_name} ({', '.join(fields)})")
        except sqlite3.OperationalError as e:
            print(f"  [WARN] index {key_name} failed: {e}")

    conn.commit()

    # ============================================================
    # 查询展示自引用元数据
    # ============================================================
    show_self_metadata(conn)

    # ============================================================
    # 验证
    # ============================================================
    show_verification(conn, csv_row_counts)

    conn.close()

    elapsed = time.time() - start_time
    print(f"\n总耗时: {elapsed:.1f} 秒")
    print("完成！")


if __name__ == "__main__":
    main()
