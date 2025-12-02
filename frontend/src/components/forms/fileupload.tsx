
export function FileForm() {
    return(
        <div className="file-form">
            <h2>Upload Document</h2>
            <input type="file" name="document" id="document" onChange={(e)=>console.log(e.target.files && e.target.files[0])}/>
        </div>
    )
}