import React, { useEffect, useState } from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Users from "./components/Users";

import UserCreate from "./components/UserCreate";
import Login from "./components/Login";


export default function App() {

  const [token, setToken] = useState<String>("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <Login />;
  }

return (

  <Router>

   <div>

   <Navbar />

   <Routes>

       <Route path="/" element={<Users />} />

       <Route path="/Schedule" element={<UserCreate />} />
       
   </Routes>

   </div>

  </Router>

);

}