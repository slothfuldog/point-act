import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';
import Login from './pages/login';
import { Routes, Route } from 'react-router-dom';
import NotFound from './pages/404';
import Private from './private-routes/privateRoute';
import Scoring from './pages/scoring';

function App() {
  const [isLogin, setIsLogin] = useState(false);
  return (
    <div>
      <Routes>
        
        <Route path='/login' element={<Login />}/>
        <Route path='/scoring' element={<Scoring />}/>
        <Route path='/*' element={<Private isLogin={isLogin} element={<NotFound/>}/>} />
      </Routes>
    </div>
  );
}

export default App;
