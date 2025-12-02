import {useState} from 'react';
import './App.css';
import { FileForm } from './components/forms/fileupload';


function App() {
    return(
        <div>
            <p>Agent Doc</p>
            <FileForm />
        </div>
    )
}

export default App
