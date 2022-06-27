import React, { useState } from "react"

export default function Teachers({ name, description, rating, teaching_category, teaching_level, teaching_subject, profilePict }) {

    return (

        <div className="card" aria-label="Post Card">
            <img src={profilePict} className="card-image-top" style={{ maxHeight: '300px', objectFit: "cover" }} />
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <p className="card-text">{teaching_subject}</p>
                <p>Rating <span className="rating">{rating}</span></p>

                <p className="deskripsi">{description}</p>
                <div className="d-flex flex-wrap">
                    <p className="card-text teaching-category bg-wow">{teaching_category}</p>
                    <p className="bg-wow">{teaching_level}</p>
                </div>
            </div>

        </div>
    )
}
