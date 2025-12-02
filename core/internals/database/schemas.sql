PRAGMA foreign_keys = ON;

-- 1. DOCUMENTS
CREATE TABLE IF NOT EXISTS documents (
    id TEXT PRIMARY KEY,           -- UUID
    filename TEXT NOT NULL,
    filepath TEXT NOT NULL,
    file_type TEXT NOT NULL CHECK(file_type IN ('pdf', 'doc', 'docx', 'txt', 'md')),
    last_read DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    total_pages INTEGER DEFAULT 0,
    total_chunks INTEGER DEFAULT 0
);

-- 2. CHUNKS 
CREATE TABLE IF NOT EXISTS document_chunks (
    id TEXT PRIMARY KEY,
    doc_id TEXT NOT NULL,
    chunk_index INTEGER,           -- Order of the chunk (0, 1, 2...)
    page_num INTEGER,              -- "See Page 5"
    content TEXT NOT NULL,         -- The actual text
    embedding BLOB,                -- The vector (JSON string or byte array)
    
    FOREIGN KEY(doc_id) REFERENCES documents(id) ON DELETE CASCADE
);

-- 3. AGENTS
CREATE TABLE IF NOT EXISTS agents (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    system_prompt TEXT NOT NULL,   -- "You are a helpful assistant..."
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 4. CONVERSATIONS
CREATE TABLE IF NOT EXISTS conversations (
    id TEXT PRIMARY KEY,
    agent_id TEXT,                 -- Which persona was used?
    doc_id TEXT,                   -- Is this chat about a specific file? (Nullable)
    title TEXT,                    -- "Chat about Resume.pdf"
    last_updated DATETIME,
    
    FOREIGN KEY(agent_id) REFERENCES agents(id),
    FOREIGN KEY(doc_id) REFERENCES documents(id) ON DELETE SET NULL
);

-- 5. MESSAGES 
CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    conversation_id TEXT NOT NULL,
    query TEXT NOT NULL,
    answer TEXT,
    message_count INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY(conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

-- 6. RAG_LOGS 
CREATE TABLE IF NOT EXISTS rag_logs (
    id TEXT PRIMARY KEY,
    conversation_id TEXT NOT NULL,
    user_query TEXT,               -- What the user asked
    retrieved_chunk_ids TEXT,      -- JSON array of chunk IDs used: "['c1', 'c2']"
    ai_response TEXT,              -- The final answer generated
    execution_time_ms INTEGER,     -- How long it took
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(conversation_id) REFERENCES conversations(id)
);
