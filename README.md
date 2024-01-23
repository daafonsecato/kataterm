# Terminal-Based Exam Application

## Overview

This application is a web-based platform designed to facilitate hands-on terminal-based exams, assessments, and interviews. It provides a user-friendly interface where candidates can receive tasks, write code, and execute commands in a terminal environment. The application is built using React and integrates a terminal emulator for real-time interaction.

## Features

- **Split View Interface**: The application features a split view interface with resizable panels. One panel displays the task details and the other provides an embedded terminal emulator.
- **Task Panel**: The left panel is dedicated to displaying task details, including instructions and requirements for the current assessment. It supports syntax highlighting and allows users to copy code snippets to the clipboard.
- **Terminal Panel**: The right panel embeds a terminal emulator connected to a server where users can execute commands and scripts. It is designed to mimic the look and feel of a real terminal.
- **Timer**: A countdown timer is displayed, providing candidates with a clear indication of the remaining time for the current task.
- **Modal Dialogs**: The application includes modal dialogs to confirm actions such as marking a task as completed. These dialogs provide additional information and prevent accidental submissions.
- **Responsive Design**: The application is styled using CSS with a focus on a clean and modern user interface that adapts to different screen sizes.
- **Accessibility**: The application is designed with accessibility in mind, ensuring that all users can navigate and interact with the platform effectively.

## Technical Details

- **Styling**: The application uses CSS for styling, with separate stylesheets for components like `App`, `TaskPanel`, and `TerminalPanel`. It features a dark theme for the terminal and contrasting colors for readability.
- **State Management**: React's `useState` hook is used to manage the application's state, including the visibility of modal dialogs and the task timer.
- **Component Architecture**: The application is composed of React components such as `TaskPanel` and `TerminalPanel`, which are modular and reusable.
- **Terminal Emulation**: The terminal panel is powered by an iframe that connects to a `ttyd` server, allowing for real-time terminal emulation.
- **Event Handling**: The application handles events such as clicks on code snippets for copying to the clipboard and toggling modal dialogs.
- **Performance**: The application includes a web vitals reporting utility to measure and optimize performance metrics.

## Usage

To use the application, candidates simply navigate to the provided URL and begin their assessment. They can read the task details, write code, and execute commands within the terminal. The application provides feedback and guidance throughout the process.

## Installation

To set up the application, clone the repository and install the necessary dependencies using a package manager like npm or yarn. You will need to configure the terminal emulator server (`ttyd`) and ensure it is running before starting the application.

## Testing

The application includes a basic test setup using `@testing-library/react` to ensure components render correctly. Additional tests can be written to cover more functionality and user interactions.

## Conclusion

This terminal-based exam application offers a streamlined environment for coding assessments and interviews, combining task management with a fully functional terminal emulator. Its intuitive design and robust feature set make it an ideal choice for evaluating technical skills in a realistic setting.