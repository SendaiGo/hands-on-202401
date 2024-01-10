-- テーブルの作成
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    passwd TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- トリガーの作成
CREATE TRIGGER IF NOT EXISTS update_user_timestamp
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- データの挿入
INSERT INTO users (name, passwd) VALUES ('sendai', 'go');

-- TODOリストのテーブル
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_completed BOOLEAN DEFAULT 0 CHECK(is_completed IN (0,1))
);

-- タイムスタンプの自動更新を設定するトリガー
CREATE TRIGGER update_timestamp
AFTER UPDATE ON todos
FOR EACH ROW 
BEGIN
    UPDATE todos SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
