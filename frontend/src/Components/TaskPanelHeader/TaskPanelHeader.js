import React from "react";
import CountdownTimer from "../CountdownTimer/CountdownTimer";
import ResetButton from "../ResetButton/ResetButton";
import ExpandButton from "../ExpandButton/ExpandButton";

const TaskPanelHeader = ({ resetQuestion }) => {
  return (
    <div className="task-panel-header">
      <h2>Task</h2>
			<CountdownTimer initialTime={'01:00:10'}/>
			<ResetButton resetQuestion={resetQuestion} />
			<ExpandButton />
    </div>
  );
};

export default TaskPanelHeader;
