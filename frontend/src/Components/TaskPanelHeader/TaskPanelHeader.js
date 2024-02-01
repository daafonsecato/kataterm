import React from "react";
import Timer from "../Timer/Timer";
import ResetButton from "../ResetButton/ResetButton";
import ExpandButton from "../ExpandButton/ExpandButton";

const TaskPanelHeader = ({ resetQuestion }) => {
  return (
    <div className="task-panel-header">
      <h2>Task</h2>
			<Timer initialTime={'00:00:00'} increment={1} />
			<ResetButton resetQuestion={resetQuestion} />
			<ExpandButton />
    </div>
  );
};

export default TaskPanelHeader;
