# DoTask

my first go program to do some task journaling

## Commands simple

    dotask list
    dotask now <title>
    dotask continue ID [now|<timestamp>]
    dotask delete ID
    dotask shutdown

    dotask ID asis|now|<timestamp> [<title> ...]

    <title> parses for:
     - asana://<id>
     - jira://<id>
     - parent://<id>

## Commands simple exmples

    dotask now admin work asana://123456789
    2014-10-10, 10:07 | 20141010-001 | admin work | asana://123456789

    dotask 001 10:05 Admin: daily dtuff ...
    2014-10-10, 10:05 | 20141010-001 | Admin: daily dtuff ... | asana://123456789

    dotask 001 asis Admin: daily stuff ...
    2014-10-10, 10:05 | 20141010-001 | Admin: daily stuff ... | asana://123456789

    dotask now asana://123456790
    2014-10-10, 10:15 | 20141010-002 | title of the asna task | asana://123456790

    dotask continue 001
    2014-10-10, 12:45 | 20141010-003 | Admin: daily stuff ... | asana://123456789

    dotask 0 now lunch
    2014-10-10, 13:05 | 20141010-004 | lunch

    dotask continue 001
    2014-10-10, 13:35 | 20141010-005 | Admin: daily stuff ... | asana://123456789

    dotask continue 002
    2014-10-10, 14:00 | 20141010-006 | title of the asna task | asana://123456790

    dotask 0 17:00 (private)
    2014-10-10, 17:00 | 20141010-007 | (private)

    dotask continue 006
    2014-10-10, 17:33 | 20141010-008 | title of the asna task | asana://123456790

    dotask 008 17:35
    2014-10-10, 17:35 | 20141010-008 | title of the asna task | asana://123456790

    dotask shutdown
    2014-10-10, 20:00 | 20141010-009 | shutdown

    dotask list
    2014-10-10, 10:05 | 20141010-001 | Admin: daily stuff ... | asana://123456789
    2014-10-10, 10:15 | 20141010-002 | title of the asna task | asana://123456790
    2014-10-10, 12:45 | 20141010-003 | Admin: daily stuff ... | asana://123456789
    2014-10-10, 13:05 | 20141010-004 | lunch
    2014-10-10, 13:35 | 20141010-005 | Admin: daily stuff ... | asana://123456789
    2014-10-10, 14:00 | 20141010-006 | title of the asna task | asana://123456790
    2014-10-10, 17:00 | 20141010-007 | (private)
    2014-10-10, 17:35 | 20141010-008 | title of the asna task | asana://123456790
    2014-10-10, 20:45 | 20141010-009 | shutdown


## Commands complex

  dotask help
  dotask --help

  dotask list[: (_today_|<date>|all)]

  dotask note: <a string: note>

  dotask shutdown

  dotask <ID> [_show_|delete|clone]

  dotask
    [<ID>|_new_|last]
    [now|(s|start)|(e|end)[:(_now_|<time>|<date+time>)]]

    [(a|asana): <an asana id string>]

    [(p|project): <a string: project name>]
    [(t|title):   <a string: title>]
    [(c|comment): <a string: comment>]


## Command complex examples

  dotask p: admin t: server upgrades
  -> dotask
      new
      start: now
      project: admin
      title: server upgrades
  <- task created:
      1408445065-7531 | 2014-08-19 12:45 [admin] server upgrades | running
      1408445065-7531 | 2014-08-19 12:45 [admin] munin check | 13:00 = 15 min



-----
