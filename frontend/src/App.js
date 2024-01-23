import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import QuestionView from "./components/QuestionView";
import StartView from "./components/StartView";
import "./App.css";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartView />} />
        <Route path="/quiz" element={<QuestionView />} />
      </Routes>
    </Router>
  );
}

export default App;
