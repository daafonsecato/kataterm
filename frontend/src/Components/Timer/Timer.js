import React, { useState, useEffect } from "react";

const Timer = ({ initialTime, increment }) => {
  
    const [initHours, initMinutes, initSeconds] = initialTime.split(":");
    const initialSeconds =
      +initHours * 60 * 60 + +initMinutes * 60 + +initSeconds;

  const [seconds, setSeconds] = useState(initialSeconds);

  useEffect(() => {
    if (seconds <= 0 && increment <= -1) {
      return;
    }
    const timer = setInterval(() => {
      setSeconds((prevSeconds) => prevSeconds + increment);
    }, 1000);
    return () => clearInterval(timer);
  }, [seconds]);

  const formatTime = (timeInSeconds) => {
    const hours = Math.floor((timeInSeconds / 60 / 60) % 24)
      .toString();
    const minutes = Math.floor((timeInSeconds / 60) % 60)
      .toString()
      .padStart(2, "0");
    const seconds = (timeInSeconds % 60).toString().padStart(2, "0");
    return `${hours > 0 ? hours + ":" : ""}${minutes}:${seconds}`;
  };

  return (
    <div className="task-timer">
      <p>{formatTime(seconds)}</p>
    </div>
  );
};
export default Timer;
