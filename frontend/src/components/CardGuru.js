// TODO: answer here
import React, { useState } from "react"


export default function CardGuru({ name, description, rating, teaching_category, teaching_level, teaching_subject }) {

    // const [name] = useState("");
    // const [description] = useState("");
    // const [rating] = useState("");
    // const [teaching_category] = useState("");
    // const [teaching_level] = useState("");
    // const [teaching_subject] = useState("");

    return (
        < div className="card col-md-3 " aria-label="Post Card">
            <h5 className="card-title">{name}</h5>
            <div className="card-body">
                <p className="card-text">{description}</p>
                <p className="card-text">{rating}</p>
                <div className="d-flex">
                    <p className="card-text">{teaching_category}</p>
                    <p className="card-text">{teaching_level}</p>
                    <p className="card-text">{teaching_subject}</p>
                </div>
            </div>
        </div >

    )
}
