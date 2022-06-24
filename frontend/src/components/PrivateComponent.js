import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import useStore from "../store/User"

const PrivateComponent = () => {
    const auth = JSON.parse(localStorage.getItem("user-info"));
    return auth ? <Outlet /> : <Navigate to={"/register"} />
}

export default PrivateComponent;