# Question
Ready lab? each 2.5s
https://manager.labs.kodekloud.com/session/readiness?lab_session_id=dfbe1ce775c54563
{"error": "not_ready"}
{result: "ready"}
## Time
GET https://c22c7370220443a4.labs.kodekloud.com/remaining_time
Resp:
{"REMAININGTIME_SECONDS":3590,"REVIEWTIME_SECONDS":600}
## Multiple Choice 4 options
GET http://backend/question
{
    "text": "What is a branch in `git`?",
    "hint": "Refer the previous lecture, Branch is nothing but a pointer to a specific commit in GIT",
    "subtext": "",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "options": [
        "git repo identifier",
        "git repo directory",
        "pointer to a specific commit in git",
        "git repo tag"
    ],
    "answer": "pointer to a specific commit in git",
    "total_questions": 16,
    "current_question_number": 1,
    "answer_statuses": [
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "text": "What is the default branch of a `git` repository?",
    "hint": "By default the branch used is `master`",
    "subtext": "",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "image": "",
    "options": [
        "release",
        "master",
        "development",
        "feature"
    ],
    "answer": "master",
    "total_questions": 16,
    "current_question_number": 2,
    "answer_statuses": [
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "text": "Sarah has been working on a git repo at `/home/sarah/story-blog` and has written a short story. Check `git log` command output in that directory to see the activity.",
    "subtext": "What's the name of the file created by Sarah?",
    "type": "multiple_choice",
    "hint": "Simply run `cd /home/sarah/story-blog; git log --name-only`",
    "staging_message": "Setting things up",
    "image": "",
    "options": [
        "story1.txt",
        "lion-and-mouse.txt",
        "frogs-and-ox.txt",
        "fox-and-grapes.txt"
    ],
    "answer": "lion-and-mouse.txt",
    "total_questions": 16,
    "current_question_number": 3,
    "answer_statuses": [
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "text": "To which branch is the :lion:  `lion-and-mouse.txt` :mouse: file committed to in the `git` repository?",
    "subtext": "",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "hint": "Check branch in the `git log --decorate` output",
    "image": "",
    "options": [
        "release",
        "master",
        "development",
        "feature"
    ],
    "answer": "master",
    "total_questions": 16,
    "current_question_number": 4,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "text": "Sarah decides to write a new story - :frog: `The Frogs and Ox` :ox:. Let's create and checkout a new branch named `story/frogs-and-ox`",
    "subtext": "",
    "type": "config_test",
    "staging_message": "Setting things up",
    "hint": "Run `git checkout -b story/frogs-and-ox`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_frogs_and_ox.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify branch",
            "test_spec_filename": "test_branch_frogs_and_ox.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "total_questions": 16,
    "current_question_number": 5,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}

{
    "text": "View the output of the `git log` command and identify the branch to which `HEAD` is pointing to now.",
    "subtext": "",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "hint": "Check branch HEAD pointer in `git log` output",
    "image": "",
    "options": [
        "release",
        "master",
        "story/frogs-and-ox",
        "sarah"
    ],
    "answer": "story/frogs-and-ox",
    "total_questions": 16,
    "current_question_number": 6,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "text": "As you can see the `HEAD` always points to the last commit on the currently checked-out branch.",
    "subtext": "",
    "type": "info",
    "staging_message": "Setting things up",
    "image": "",
    "options": [
        "Ok"
    ],
    "answer": "Ok",
    "total_questions": 16,
    "current_question_number": 7,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "text": "Sarah is half way through the :frog: Frogs and Ox :ox: story. It's not complete yet.",
    "subtext": "View the story she has written in the file `frogs-and-ox.txt`",
    "type": "info",
    "staging_image": "/images/emojis/sarah-working.png",
    "staging_image_border": "unset",
    "staging_image_width": "120px",
    "staging_message": "Sarah is writing the `Frogs and Ox` story....",
    "options": [
        "Ok"
    ],
    "answer": "Ok",
    "before": [
        {
            "type": "command",
            "command": "sleep 5; docker exec dev01 sh -c 'cp /tmp/stories/frogs-and-ox-half.txt /home/sarah/story-blog/frogs-and-ox.txt'",
            "shell": true
        }
    ],
    "total_questions": 16,
    "current_question_number": 8,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ],
    "hasBeforeAction": true
}
## Config_test Question
{
    "text": "Sarah decides to write a new story - :frog: `The Frogs and Ox` :ox:. Let's create and checkout a new branch named `story/frogs-and-ox`",
    "subtext": "",
    "type": "config_test",
    "staging_message": "Setting things up",
    "hint": "Run `git checkout -b story/frogs-and-ox`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_frogs_and_ox.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify branch",
            "test_spec_filename": "test_branch_frogs_and_ox.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "total_questions": 16,
    "current_question_number": 5,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
## Info Question
{
    "text": "Access the web application using the tab 'simple-webapp-ui' above the terminal window.",
    "type": "info",
    "options": [
        "Ok"
    ],
    "answer": "Ok",
    "links": [
        {
            "type": "port",
            "name": "simple-webapp-ui",
            "port": "30080"
        }
    ],
    "total_questions": 11,
    "current_question_number": 11,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current"
    ]
}
# Actions



## Skip question
GET http://backend/skip_question
    {
        "total_questions": 11,
        "answer_statuses": [
            "completed",
            "completed",
            null,
            null,
            null,
            null,
            null,
            null,
            null,
            null,
            null
        ],
        "new_question_number": 3
}
## Config after some questions (Purpose TBD)
GET http://backend/config
Sometimes this answers an empty response
{
    "environment": "Prod",
    "quiz_mode": "standard",
    "review_question": false,
    "exam_evaluation_mode": "SHOW_RESULT_AND_SCORE",
    "under_review": false,
    "run_tests_one_at_a_time": false,
    "show_cmd_output_on_error": false,
    "layout": "horizontal",
    "start_from": 8,
    "panel_first": false
}
## Heart beat (Each 20 seconds)
GET https://backend/heart_beat
200 or non response status code

## Submit Answer
POST https://7d78542a085e4798.labs.kodekloud.com/submit_answer
Payload: {answer: "6443"}
200 Accepted
or
400 "Incorrect! Try Again."
## In case of stage before action step
GET http://backend/question
{
    "text": "Sarah is half way through the :frog: Frogs and Ox :ox: story. It's not complete yet.",
    "subtext": "View the story she has written in the file `frogs-and-ox.txt`",
    "type": "info",
    "staging_image": "/images/emojis/sarah-working.png",
    "staging_image_border": "unset",
    "staging_image_width": "120px",
    "staging_message": "Sarah is writing the `Frogs and Ox` story....",
    "options": [
        "Ok"
    ],
    "answer": "Ok",
    "before": [
        {
            "type": "command",
            "command": "sleep 5; docker exec dev01 sh -c 'cp /tmp/stories/frogs-and-ox-half.txt /home/sarah/story-blog/frogs-and-ox.txt'",
            "shell": true
        }
    ],
    "total_questions": 16,
    "current_question_number": 8,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ],
    "hasBeforeAction": true
}
### Post stage before actions
POST https://backend/stage_before_actions
PAYLOAD {question_number: 8}
{
    "text": "Sarah is half way through the :frog: Frogs and Ox :ox: story. It's not complete yet.",
    "subtext": "View the story she has written in the file `frogs-and-ox.txt`",
    "type": "info",
    "staging_image": "/images/emojis/sarah-working.png",
    "staging_image_border": "unset",
    "staging_image_width": "120px",
    "staging_message": "Sarah is writing the `Frogs and Ox` story....",
    "options": [
        "Ok"
    ],
    "answer": "Ok",
    "before": [
        {
            "type": "command",
            "command": "sleep 5; docker exec dev01 sh -c 'cp /tmp/stories/frogs-and-ox-half.txt /home/sarah/story-blog/frogs-and-ox.txt'",
            "shell": true
        }
    ],
    "hasBeforeAction": false,
    "total_questions": 16,
    "current_question_number": 8,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}

 cat frogs-and-ox.txt 
--------------------------------------------
      THE FROGS AND THE OX
--------------------------------------------

An Ox came down to a reedy pool to drink. As he splashed heavily into the water, he crushed a young Frog into the mud.

The old Frog soon missed the little one and asked his brothers and sisters what had become of him.

"A great big monster," said one of them, "stepped on little brother with one of his huge feet!"

"Big, was he!" said the old Frog, puffing herself up. "Was he as big as this?"

"Oh, much ....


backend/question
{
    "before": [
        {
            "type": "command",
            "command": "sleep 5",
            "shell": true
        }
    ],
    "text": "Max informs Sarah that in her first story there's a typo in the title and needs to be fixed ASAP!",
    "subtext": "We must go back and fix the story in the `master` branch. But before we do that, let's commit the new story we have written so far. We don't want to carry our incomplete story to the master branch. \n\n Stage and commit the story with the message `Add incomplete frogs-and-ox story` ",
    "staging_image": "/images/emojis/max.png",
    "staging_image_border": "unset",
    "staging_image_width": "120px",
    "staging_message": "The phone :phone: rings! It's Max!",
    "type": "config_test",
    "hint": "Run `git add frogs-and-ox.txt; git commit -am 'Add incomplete frogs-and-ox story'`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_sarah_commit.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify commit",
            "test_spec_filename": "test_branch_sarah_commit.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "total_questions": 16,
    "current_question_number": 9,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ],
    "hasBeforeAction": true
}
https://c22c7370220443a4.labs.kodekloud.com/stage_before_actions
{question_number: 9}
{
    "before": [
        {
            "type": "command",
            "command": "sleep 5",
            "shell": true
        }
    ],
    "text": "Max informs Sarah that in her first story there's a typo in the title and needs to be fixed ASAP!",
    "subtext": "We must go back and fix the story in the `master` branch. But before we do that, let's commit the new story we have written so far. We don't want to carry our incomplete story to the master branch. \n\n Stage and commit the story with the message `Add incomplete frogs-and-ox story` ",
    "staging_image": "/images/emojis/max.png",
    "staging_image_border": "unset",
    "staging_image_width": "120px",
    "staging_message": "The phone :phone: rings! It's Max!",
    "type": "config_test",
    "hint": "Run `git add frogs-and-ox.txt; git commit -am 'Add incomplete frogs-and-ox story'`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_sarah_commit.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify commit",
            "test_spec_filename": "test_branch_sarah_commit.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "hasBeforeAction": false,
    "total_questions": 16,
    "current_question_number": 9,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
POST https://c22c7370220443a4.labs.kodekloud.com/check_config
{configNumber: {isTrusted: true, _vts: 1704502442186}}
{
    "message": "Tasks completed.",
    "additional_details": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_sarah_commit.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify commit",
            "test_spec_filename": "test_branch_sarah_commit.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": "",
            "test_result": true,
            "cmd_output": "============================= test session starts ==============================\nplatform linux -- Python 3.9.7, pytest-7.0.1, pluggy-1.0.0\nrootdir: /questions\nplugins: testinfra-6.6.0\ncollected 1 item\n\ntest_branch_sarah_commit.py .                                            [100%]\n\n=============================== warnings summary ===============================\ntest_branch_sarah_commit.py:9\n  /questions/test_branch_sarah_commit.py:9: DeprecationWarning: invalid escape sequence \\*\n    cmd = host.run(\"cd /home/sarah/story-blog;git show-branch -a | grep '\\*' | grep -v `git rev-parse --abbrev-ref HEAD` |head -n1\")\n\n-- Docs: https://docs.pytest.org/en/stable/how-to/capture-warnings.html\n========================= 1 passed, 1 warning in 1.29s =========================\n"
        }
    ]
}

backend/question/
{
    "text": "Now checkout the `master` branch.",
    "subtext": "",
    "type": "config_test",
    "hint": "Run `git checkout master`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_master.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify branch",
            "test_spec_filename": "test_branch_master.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "total_questions": 16,
    "current_question_number": 10,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
{
    "message": "Tasks completed.",
    "additional_details": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_master.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify branch",
            "test_spec_filename": "test_branch_master.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": "",
            "test_result": true,
            "cmd_output": "============================= test session starts ==============================\nplatform linux -- Python 3.9.7, pytest-7.0.1, pluggy-1.0.0\nrootdir: /questions\nplugins: testinfra-6.6.0\ncollected 1 item\n\ntest_branch_master.py .                                                  [100%]\n\n============================== 1 passed in 0.15s ===============================\n"
        }
    ]
}
{
    "text": "Let's fix the typo in the `lion-and-mouse.txt` file. `LION`  :lion:  is mis-spelt as `LIOON`. Please fix it and then commit the changes.",
    "subtext": "Commit message: `Fix typo in story title`",
    "type": "config_test",
    "staging_message": "Setting things up",
    "hint": "Use vi editor to edit the file and fix the typo. Then run the command `git commit -am 'Fix typo in story title'`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_master_fix.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify fix and commit",
            "test_spec_filename": "test_branch_master_fix.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "total_questions": 16,
    "current_question_number": 11,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null
    ]
}

## Check config_test question (not completed)
backend/check_config/
{
    "message": "Tasks not completed!",
    "additional_details": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_master_fix.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify fix and commit",
            "test_spec_filename": "test_branch_master_fix.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": "",
            "test_result": false,
            "err_message": "Command 'docker exec -w /questions test_infra pytest test_branch_master_fix.py' returned non-zero exit status 1.",
            "cmd_output": "============================= test session starts ==============================\nplatform linux -- Python 3.9.7, pytest-7.0.1, pluggy-1.0.0\nrootdir: /questions\nplugins: testinfra-6.6.0\ncollected 1 item\n\ntest_branch_master_fix.py F                                              [100%]\n\n=================================== FAILURES ===================================\n_______________________________ test_git_install _______________________________\n\n    def test_git_install():\n        host = testinfra.get_host(\"docker://dev01\")\n        branch = \"master\"\n        cmd = host.run(\"cd /home/sarah/story-blog;git symbolic-ref --short -q HEAD\")\n        assert branch == cmd.stdout.rstrip(),  \"In '/home/sarah/story-blog' active branch is not %s\" %branch\n    \n        cmd = host.run(\"cat /home/sarah/story-blog/lion-and-mouse.txt\")\n>       assert \"THE LION AND THE MOUSE\" in cmd.stdout,  \" - Typo in 'lion-and-mouse.txt' is not fixed\"\nE       AssertionError:  - Typo in 'lion-and-mouse.txt' is not fixed\nE       assert 'THE LION AND THE MOUSE' in '--------------------------------------------\\n      THE LIOON AND THE MOUSE\\n----------------------------------------...free.\\n\\n\"You laughed when I said I would repay you,\" said the Mouse. \"Now you see that even a Mouse can help a Lion.\"'\nE        +  where '--------------------------------------------\\n      THE LIOON AND THE MOUSE\\n----------------------------------------...free.\\n\\n\"You laughed when I said I would repay you,\" said the Mouse. \"Now you see that even a Mouse can help a Lion.\"' = CommandResult(command=b'cat /home/sarah/story-blog/lion-and-mouse.txt', exit_status=0, stdout=b'----------------------...laughed when I said I would repay you,\" said the Mouse. \"Now you see that even a Mouse can help a Lion.\"', stderr=None).stdout\n\ntest_branch_master_fix.py:10: AssertionError\n=========================== short test summary info ============================\nFAILED test_branch_master_fix.py::test_git_install - AssertionError:  - Typo ...\n============================== 1 failed in 0.34s ===============================\n"
        }
    ]
}

## Check config_test question (completed)
backend/check_config/
{
    "message": "Tasks completed.",
    "additional_details": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_master_fix.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify fix and commit",
            "test_spec_filename": "test_branch_master_fix.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": "",
            "test_result": true,
            "cmd_output": "============================= test session starts ==============================\nplatform linux -- Python 3.9.7, pytest-7.0.1, pluggy-1.0.0\nrootdir: /questions\nplugins: testinfra-6.6.0\ncollected 1 item\n\ntest_branch_master_fix.py .                                              [100%]\n\n============================== 1 passed in 0.56s ===============================\n"
        }
    ]
}
{
    "text": "Sarah has now finished her story. Check the changes and commit them with the message `Completed frogs-and-ox story`",
    "subtext": "",
    "type": "config_test",
    "staging_message": "Sarah is finishing her story",
    "hint": "Run the command `git commit -am 'Completed frogs-and-ox story'`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_sarah_final_commit.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify branch",
            "test_spec_filename": "test_branch_sarah_final_commit.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "before": [
        {
            "type": "command",
            "command": "sleep 5; docker exec dev01 sh -c 'cp /tmp/stories/frogs-and-ox.txt /home/sarah/story-blog/frogs-and-ox.txt'",
            "shell": true
        }
    ],
    "total_questions": 16,
    "current_question_number": 13,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null
    ],
    "hasBeforeAction": true
}
{
    "text": "Sarah has now finished her story. Check the changes and commit them with the message `Completed frogs-and-ox story`",
    "subtext": "",
    "type": "config_test",
    "staging_message": "Sarah is finishing her story",
    "hint": "Run the command `git commit -am 'Completed frogs-and-ox story'`",
    "image": "",
    "tests": [
        {
            "command": "docker exec -w /questions test_infra pytest test_branch_sarah_final_commit.py",
            "name": "dist",
            "shell": true,
            "spec": "Verify branch",
            "test_spec_filename": "test_branch_sarah_final_commit.py",
            "type": "testinfra",
            "user_executed": false,
            "cmd_prefix": ""
        }
    ],
    "before": [
        {
            "type": "command",
            "command": "sleep 5; docker exec dev01 sh -c 'cp /tmp/stories/frogs-and-ox.txt /home/sarah/story-blog/frogs-and-ox.txt'",
            "shell": true
        }
    ],
    "hasBeforeAction": false,
    "total_questions": 16,
    "current_question_number": 13,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null
    ]
}
{
    "environment": "Prod",
    "quiz_mode": "standard",
    "review_question": false,
    "exam_evaluation_mode": "SHOW_RESULT_AND_SCORE",
    "under_review": false,
    "run_tests_one_at_a_time": false,
    "show_cmd_output_on_error": false,
    "layout": "horizontal",
    "start_from": 13,
    "panel_first": false
}
{
    "text": "A new git repository is created at the path `/home/sarah/website` for hosting the story website.",
    "subtext": "Count the number of branches available in that repository including the `master` branch.",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "hint": "Change to directory `cd /home/sarah/website` and run `git branch` command",
    "image": "",
    "options": [
        "1",
        "2",
        "3",
        "4",
        "5"
    ],
    "answer": "5",
    "total_questions": 16,
    "current_question_number": 14,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null
    ]
}
{
    "text": "A new git repository is created at the path `/home/sarah/website` for hosting the story website.",
    "subtext": "Count the number of branches available in that repository including the `master` branch.",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "hint": "Change to directory `cd /home/sarah/website` and run `git branch` command",
    "image": "",
    "options": [
        "1",
        "2",
        "3",
        "4",
        "5"
    ],
    "answer": "5",
    "total_questions": 16,
    "current_question_number": 14,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null
    ]
}
{
    "text": "Looking at the commit history, try to guess what branch was the `feature/signout` branch created from?",
    "subtext": "Checkout branch `feature/signout` and then use the command `git log --graph --decorate` to see previous commit history along with the branch they were committed on.",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "hint": "Checkout to directory `git checkout feature/signout; git log --graph --decorate`",
    "image": "",
    "options": [
        "master",
        "feature/signup",
        "feature/signout",
        "feature/cart",
        "feature/checkout"
    ],
    "answer": "feature/signup",
    "total_questions": 16,
    "current_question_number": 15,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null
    ]
}
{
    "text": "Here's a fun one! Looking at the commit history of all branches, what's the best graphical representation of the branches in this repository?",
    "subtext": "Checkout each branch and then use the command `git log --graph --decorate` to see previous branch. \n\n A. ![branch-1](https://res.cloudinary.com/cloudusthad/image/upload/w_300,c_scale/v1597113199/git/git-branch-2.png) \n\n  B. ![branch-1](https://res.cloudinary.com/cloudusthad/image/upload/w_400,c_scale/v1597113200/git/git-branch-4.png) \n\n C. ![branch-1](https://res.cloudinary.com/cloudusthad/image/upload/w_400,c_scale/v1597113432/git/git-branch-1.png) \n\n D. ![branch-1](https://res.cloudinary.com/cloudusthad/image/upload/w_400,c_scale/v1597113199/git/git-branch-3.png)",
    "type": "multiple_choice",
    "staging_message": "Setting things up",
    "hint": "Checkout to directory `git checkout feature/signout; git log --graph --decorate`",
    "image": "",
    "options": [
        "A",
        "B",
        "C",
        "D"
    ],
    "answer": "D",
    "total_questions": 16,
    "current_question_number": 16,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current"
    ]
}