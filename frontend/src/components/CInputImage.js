import React, { useEffect, useState } from 'react'
/*
* onFileChange => return callback with base64 string of image
*/
function CInputImage({ onFileChange = (base64String) => { }, value, isShowUploadButton, resetDefaultState }) {
    const [preview, setPreview] = useState(value)
    const _handleFileChange = (e) => {
        e.preventDefault()
        const objectUrl = URL.createObjectURL(e.target.files?.[0])

        //convert image to base64 format
        var reader = new FileReader();
        reader.onloadend = function () {
            //   console.log('Base64 string', reader.result)
            const base64String = reader.result.split("base64,")?.[1]
            onFileChange(base64String)
        }
        reader.readAsDataURL(e.target.files?.[0])

        setPreview(objectUrl)
    }

    useEffect(() => {
        if (resetDefaultState) {
            setPreview(value)
        }

    }, [resetDefaultState])


    const display = isShowUploadButton ? 'inline' : 'none'
    return (
        <div>
            <img src={preview} style={{ maxHeight: '300px', objectFit: 'cover' }} />
            <input style={{ display: display }} type="file" onChange={_handleFileChange} />
        </div>
    )
}

export default CInputImage