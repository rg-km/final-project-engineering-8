import React, { useEffect } from 'react';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './App.css'
import { Routes, Route } from 'react-router-dom'
import Login from './components/Login'
import Register from './components/Register'
import RegisterTeacher from './components/RegisterTeacher'
import Home from './components/Home';
import TeacherDetail from './components/TeacherDetail';
import StudentProfile from './components/StudentProfile'
import PrivateComponent from './components/PrivateComponent'
import LoginTeacher from './components/LoginTeacher'
import TeacherProfile from './components/TeacherProfile'
import NotFound from './components/NotFound';

function App() {

  useEffect(() => {
    document.title = `HALLOGURU`;
  });

  return (

    <div className="App">

      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route exact path="home" element={<Home />} />
        <Route path="login" element={<Login />} />
        <Route path="register" element={<Register />} />
        <Route path="register/teacher" element={<RegisterTeacher />} />
        <Route path="login/teacher" element={<LoginTeacher />} />
        <Route element={<PrivateComponent />}>
          <Route path="teacher/:id" element={<TeacherDetail />} />
          <Route path="profile" element={<StudentProfile />} />
          <Route path="profile/teacher/:id" element={<TeacherProfile />} />
        </Route>
        <Route path="*" element={<NotFound />} />
      </Routes>

    </div >

  )
}
export default App