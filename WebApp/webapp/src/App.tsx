import React from 'react';
import logo from './logo.svg';
import './App.css';

import Home from './pages/Home';
import Matchups from './pages/Matchups';

import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import NavBar from './components/NavBar';

function App() {
  return (
    <div>
      <Router>
        <NavBar/>
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/matchups" element={<Matchups/>} />
          <Route path="/predictions" element={<Matchups/>} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
