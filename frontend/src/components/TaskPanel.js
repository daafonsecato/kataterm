// TaskPanel.js
import React, { useEffect, useState, useRef } from 'react';
import './TaskPanel.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faExpand, faUndo } from '@fortawesome/free-solid-svg-icons';
import { marked } from 'marked';

function TaskPanel({ onCheckClick }) {
    const [taskDetails, setTaskDetails] = useState('');
    const [timeLeft, setTimeLeft] = useState('59:59');
    const [isLoading, setIsLoading] = useState(false); // Add isLoading state
    const taskPanelRef = useRef(null);

    useEffect(() => {
        const timerInterval = setInterval(() => {
            const [minutes, seconds] = timeLeft.split(':').map(Number);

            if (minutes === 0 && seconds === 0) {
                clearInterval(timerInterval);
            } else if (seconds === 0) {
                setTimeLeft(`${minutes - 1}:${59}`);
            } else {
                if (seconds < 11) {
                    console.log('Decreasing seconds');
                    setTimeLeft(`${minutes}:0${seconds - 1}`);
                } else {
                    setTimeLeft(`${minutes}:${seconds - 1}`);
                }
            }
        }, 1000);

        return () => {
            clearInterval(timerInterval);
        };
    }, [timeLeft]);


    const fetchTaskDetails = async () => {
        try {
            const response = await fetch('http://terminal.kataterm.com:8000/question');
            let data;
            data = await response.json();
            console.log(data);
            if (response.status == 200) {
                if (data.before_actions.length > 0) {
                    setupQuestion(data)
                } else {
                    setIsLoading(false); // Hide loader
                }
            }
            setTaskDetails(data);
        } catch (error) {
            console.error('Error fetching question:', error);
        }
    };

    async function setupQuestion(data) {
        try {
            setIsLoading(true); // Show loader
            console.log("is loading");
            const response = await fetch('http://terminal.kataterm.com:8000/stage_before_actions', {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "ID": data.id,
                    "before_actions": data.before_actions
                })
            });
            data = await response.json();
            console.log(data);
            setTaskDetails(data);
        } catch (error) {
            console.error('Error submitting answer:', error);
        } finally {
            setIsLoading(false); // Hide loader
        }
    }

    const isMountedRef = useRef(false);

    useEffect(() => {
        if (!isMountedRef.current) {
            fetchTaskDetails();
            isMountedRef.current = true;
        }
    }, []);

    async function checkConfig(questionID) {
        try {
            const response = await fetch('http://terminal.kataterm.com:8000/check_config', {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "ID": questionID
                })
            });
            const status = await response.status;

            if (status == 200) {
                console.log('Answer submitted:', status);
                fetchTaskDetails();
            }
        } catch (error) {
            console.error('Error submitting answer:', error);
        }
    }

    async function submitAnswer(answer) {
        try {
            const response = await fetch('http://terminal.kataterm.com:8000/submit_answer', {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ "Answer": answer })
            });
            const status = await response.status;

            if (status == 200) {
                console.log('Answer submitted:', answer);
                fetchTaskDetails();
            }
        } catch (error) {
            console.error('Error submitting answer:', error);
        }
    }

    function transformJsonToComponent(jsonData) {
        if (jsonData.type === 'multiple_choice') {
            return (
                <div className="flexible-container">
                    <div className="navigation">
                        <div className="progress-bar" style={{ display: 'none' }}></div>
                        <div className="question-numbers">
                            {jsonData.current_question_number}/{jsonData.total_questions}
                        </div>
                    </div>
                    <div className="main-content">
                        <div className="question-text">{jsonData.text}</div>
                        <div className="answer-buttons">
                            {jsonData.options.map((option, index) => {
                                const isCorrect = option === jsonData.answer;
                                const answerStatus =
                                    jsonData.answer_statuses[index] === 'current' ? 'current' : '';

                                return (
                                    <button
                                        key={index}
                                        className={`answer-button ${answerStatus}`}
                                        onClick={() => submitAnswer(option)}
                                    >
                                        {option}
                                        <span className={`answer-icon ${isCorrect ? 'correct' : 'incorrect'}`}></span>
                                    </button>
                                );
                            })}
                        </div>
                    </div>
                </div>
            );
        } else if (jsonData.type === 'config_test') {
            return (
                <div className="flexible-container">
                    <div className="navigation">
                        <div className="progress-bar" style={{ display: 'none' }}></div>
                        <div className="question-numbers">
                            {jsonData.current_question_number}/{jsonData.total_questions}
                        </div>
                    </div>
                    <div className="main-content">
                        <div className="question-text" dangerouslySetInnerHTML={{ __html: marked(jsonData.text) }}></div>
                        <button
                            className="check-button"
                            onClick={() => checkConfig(jsonData.current_question_number)}
                        >
                            Check
                        </button>
                        <div className="test-list">
                            {jsonData.tests && jsonData.tests.map((test, index) => {
                                const testStatus = test.user_executed
                                    ? test.status === 'passed'
                                        ? 'passed'
                                        : 'failed'
                                    : 'idle';

                                return (
                                    <div key={index} className="test-item">
                                        <span className={`test-status ${testStatus}`}></span>
                                        <span className="test-name">{test.spec}</span>
                                    </div>
                                );
                            })}
                        </div>
                    </div>
                </div>
            );
        }

        return null;
    }

    const resetQuestion = async (event) => {
        try {
            setupQuestion(taskDetails);
        } catch (error) {
            console.error('Error resetting answer:', error);
        }
    }

    const handlePanelClick = async (event) => {
        if (event.target.tagName === 'CODE') {
            try {
                await navigator.clipboard.writeText(event.target.textContent);
            } catch (err) {
                console.error('Unable to copy code:', err);
            }
        }
    };

    const skipQuestion = async (event) => {
        try {
            const response = await fetch('http://terminal.kataterm.com:8000/skip_question', {
                method: 'GET',
            });
            const data = await response.json();
            console.log('Answer submitted:', data);
        } catch (error) {
            console.error('Error submitting answer:', error);
        }
    }

    return (
        <div className="TaskPanel" ref={taskPanelRef} onClick={handlePanelClick}>
            <div className="task-panel-header">
                <h2>Task</h2>
                <div className="task-timer">{timeLeft}</div>
                <button className="reset-button" onClick={resetQuestion}>
                    <FontAwesomeIcon icon={faUndo} />
                </button>
                <button className="expand-button">
                    <FontAwesomeIcon icon={faExpand} />
                </button>
            </div>
            {isLoading ? ( // Conditional rendering for loader
                <div className="loader">Loading...</div>
            ) : (
                transformJsonToComponent(taskDetails)
            )}

            <div className="task-panel-footer">
                <button className="check-button" onClick={onCheckClick}>Check</button>
                <button className="skip-button" onClick={skipQuestion}>Skip</button>
                <button className="try-later-button">Try Later</button>
            </div>
        </div>
    );
}

export default TaskPanel;
