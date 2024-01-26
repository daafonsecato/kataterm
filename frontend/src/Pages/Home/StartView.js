import React from 'react';
import { useNavigate } from 'react-router-dom';

const StartView = () => {
    const navigate = useNavigate();

    const handleStart = () => {
        navigate('/quiz');
    };

    return (
        <div>
            <h1>Welcome to Quizard!</h1>
            <button onClick={handleStart}>Start</button>
        </div>
    );
};

export default StartView;
