import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExpand } from "@fortawesome/free-solid-svg-icons";

const ExpandButton = () => {
  return (
    <button className="expand-button">
      <FontAwesomeIcon icon={faExpand} />
    </button>
  );
};

export default ExpandButton;
