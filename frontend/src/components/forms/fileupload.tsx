import React from"react";
import { useState } from "react";

export function FileForm() {
    const [text,setText] = React.useState<string>("No Drop Yet");
    return(
        <div className="file-form">
            <h2>Upload Document</h2>
            <input className="bg-black text-white" type="file" name="document" id="document" onChange={(e)=>console.log(e.target.files && e.target.files[0])}/>
            <p>{text}</p>
        </div>
    )
}