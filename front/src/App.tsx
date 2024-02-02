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
  const [isLogin, setIsLogin] = useState(false);
  const [role, setRole] = useState(5);

  const keepLogin = () =>{
    let getLocaleStorage = localStorage.getItem("p-log-43");
    if(getLocaleStorage){
      let res = Axios.post("localhost:3000/" + "/keep-login",{});

    }
  }

  useEffect(() =>{
    keepLogin();
  }, [])
  return (
    <div>
      <Routes>
        
        <Route path='/login' element={<Login />}/>
        <Route path='/scoring' element={<Private isLogin={isLogin} element={<Scoring role={role} />}/>} />
        <Route path='/*' element={<Private isLogin={isLogin} element={<NotFound/>}/>} />
      </Routes>
    </div>
  );
}

export default App;
