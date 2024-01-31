import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUndo } from "@fortawesome/free-solid-svg-icons";

const ResetButton = ({resetQuestion}) => {
  return (
    <button className="reset-button" onClick={resetQuestion}>
      <FontAwesomeIcon icon={faUndo} />
    </button>
  );
};

export default ResetButton;
