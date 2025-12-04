package database

const(
	AddDocumentQ = `
	INSERT INTO documents (id,filename,filepath,file_type,created_at,last_read,total_pages,total_chunks)
		VALUES (?,?,?,?,COALESCE(?, STRFTIME('%Y-%m-%d %H:%M:%S', 'now')), COALESCE(?, STRFTIME('%Y-%m-%d %H:%M:%S', 'now')),?,?)
		ON CONFLICT(id) DO UPDATE SET
			filename = excluded.filename,
			filepath = excluded.filepath,
			file_type = excluded.file_type,
			total_pages = excluded.total_pages,
			total_chunks = excluded.total_chunks
		RETURNING id, created_at;`
	Get_DocumentsAllQ=`SELECT * FROM documents;`
	Get_DocumentsByPathQ=`SELECT * FROM documents WHERE filepath=$1;`
)