import React, { useCallback, useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import Login from './pages/login';
import { Routes, Route } from 'react-router-dom';
import NotFound from './pages/404';
import Private from './private-routes/privateRoute';
import Scoring from './pages/scoring';
import Axios from "axios";

function App() {
  const getLocaleStorage = JSON.parse(localStorage.getItem("login") || '{}');
  const [isLogin, setIsLogin] = useState(getLocaleStorage.isLogin ? true : false);
  const [role, setRole] = useState(5);

  

  useEffect(() =>{
    const keepLogin = async () =>{
      try{
      console.log(getLocaleStorage.username)
      console.log("console",isLogin)
      if(getLocaleStorage){
        let res =  await Axios.post("http://localhost:3000/" + "keep-login",{
          username: getLocaleStorage.username,
          isLogin: getLocaleStorage.isLogin
        })
        localStorage.setItem("login", JSON.stringify(res.data.response))
        setIsLogin(true)
        console.log(isLogin)
        setRole(res.data.response.role)
        console.log(role)
      }
    }catch (e){
      setIsLogin(false)
      localStorage.removeItem("login")
      console.log(e)
    }
    }
    keepLogin()
  }, [isLogin, getLocaleStorage, role])
  return (
    <div>
      <Routes>
        
        <Route path='/login' element={<Login isLogin={isLogin}/>}/>
        <Route path='/scoring' element={<Private isLogin={isLogin} element={<Scoring role={role} />}/>} />
        <Route path='/*' element={<Private isLogin={isLogin} element={<NotFound/>}/>} />
      </Routes>
    </div>
  );
}

export default App;
