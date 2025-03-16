import { Route,Routes } from "react-router-dom";
import React from "react";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import ProfilePage from "./pages/ProfilePage";
import CreatNewAccount from "./pages/CreateNewAccount"

function App(){
  return (<div>
    <Routes>
    <Route path="/" element={<LoginPage/>}></Route>
    <Route path="/register" element={<CreatNewAccount/>}></Route>
    <Route path="/home" element={<HomePage/>}/>
    <Route path="/progile" element={<ProfilePage/>}/>
  </Routes>
  </div>
  
  );
}
export  default App;