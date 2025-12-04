import { useState, useEffect } from 'react';
import './App.css';
// @ts-ignore
import { SelectFile } from "../wailsjs/go/main/App";
// @ts-ignore
import { EventsOn } from "../wailsjs/runtime/runtime";

function App() {
    const [files, setFiles] = useState<string[]>([]);
    const [isDragging, setIsDragging] = useState(false);

    useEffect(() => {
        const fileListener= EventsOn("drop", (paths: string[]) => {            
            if (paths && Array.isArray(paths) && paths.length > 0) {
                setFiles((prev) => [...new Set([...prev, ...paths])]);
                setIsDragging(false);
            }
        });

        // Cleanup: This function runs when the component unmounts
        return () => {}
    }, []);
    // 2. GLOBAL DRAG & DROP HANDLER (The Visuals)
    useEffect(() => {
        // We attach to window to catch drops ANYWHERE in the app
        const handleDragEnter = (e: DragEvent) => {
            e.preventDefault();
            setIsDragging(true);
        };

        const handleDragOver = (e: DragEvent) => {
            e.preventDefault(); // Crucial: allows the drop to happen
            setIsDragging(true);
        };

        const handleDragLeave = (e: DragEvent) => {
            e.preventDefault();
            // Only hide if the mouse actually leaves the window
            if (e.clientX === 0 && e.clientY === 0) {
                setIsDragging(false);
            }
        };

        const handleDrop = (e: DragEvent) => {
            e.preventDefault(); // Stop browser from opening file
            setIsDragging(false); // Hide overlay
        };

        // Add Listeners
        window.addEventListener('dragenter', handleDragEnter);
        window.addEventListener('dragover', handleDragOver);
        window.addEventListener('dragleave', handleDragLeave);
        window.addEventListener('drop', handleDrop);

        // Cleanup
        return () => {
            window.removeEventListener('dragenter', handleDragEnter);
            window.removeEventListener('dragover', handleDragOver);
            window.removeEventListener('dragleave', handleDragLeave);
            window.removeEventListener('drop', handleDrop);
        };
    }, []);

    // 3. MANUAL SELECTION
    const handleManualSelect = async () => {
        try {
            const path = await SelectFile();
            if (path){
                setFiles((prev) => [...new Set([...prev, path])]);
            }

        } catch (err) {
            console.error(err);
        }
    };

    return (
        <div className="container">
            {/* Overlay: Only shows when dragging */}
            <div className={`drag-overlay ${isDragging ? 'active' : ''}`}>
                <div className="drag-message">Drop Files Here</div>
            </div>

            <div className="card">
                <h1>File Importer</h1>
                <p style={{ color: '#666' }}>
                    Drag files anywhere or click below
                </p>

                <button className="btn-select" onClick={handleManualSelect}>
                    Select File Manually
                </button>

                {files.length > 0 && (
                    <div className="file-list">
                        <strong>Captured Paths:</strong>
                        {files.map((f, i) => (
                            <div key={i} className="file-item">{f}</div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
}

export default App;