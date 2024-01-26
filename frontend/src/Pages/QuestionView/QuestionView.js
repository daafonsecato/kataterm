import React from 'react';
import TaskPanel from '../../Components/TaskPanel/TaskPanel';
import TerminalPanel from '../../Components/TerminalPanel/TerminalPanel';
import Split from 'react-split';
import EditorPanel from '../../Components/EditorPanel/EditorPanel';

function QuestionView() {
    const ttydUrl = process.env.REACT_APP_TTYD_SERVER_URL; // Replace with your ttyd URL
    const codeServerUrl = process.env.REACT_APP_CODE_SERVER_URL; // Replace with your ttyd URL

    return (
        <div className="QuestionView">
            <div className="BasePanels">
                <Split
                    className="App"
                    direction="horizontal" // New split direction
                    sizes={[30, 70]} // Initial sizes of the three panels (left, middle, right)
                    minSize={100} // Minimum size of each panel
                    expandToMin={false} // Prevent panels from collapsing to minimum size
                    gutterSize={10} // Size of the gutter between panels
                    gutterAlign="center" // Align the gutter in the center
                    snapOffset={30} // Snap to minimum                                                   size if within 30 pixels
                >
                    <div className="LeftPanel">
                        <TaskPanel />
                    </div>

                    <Split
                        className="RightPanel"
                        direction="vertical" // New split direction
                        sizes={[60, 40]} // Initial sizes of the two panels (top, bottom)
                        minSize={10} // Minimum size of each panel
                        expandToMin={false} // Prevent panels from collapsing to minimum size
                        gutterSize={10} // Size of the gutter between panels
                        gutterAlign="center" // Align the gutter in the center
                        snapOffset={30} // Snap to minimum size if within 30 pixels
                    >
                        <div className="TopPanel">
                            <EditorPanel codeServerUrl={codeServerUrl} />
                        </div>
                        <div className="BottomPanel">
                            <TerminalPanel wsUrl={ttydUrl} />
                        </div>
                    </Split>
                </Split>
            </div>

        </div>
    );
}

export default QuestionView;