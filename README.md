# DoTask

my first go program to do some task journaling

## Commands

    dotask (l)ist | (s)how
    dotask (n)ow <title>
    dotask (c)ontinue <ID> [now|<timestamp>]
    dotask delete <ID>
    dotask shutdown

    dotask <ID> asis|now|<timestamp> [<title> ...]

    <title> parses for:
     - asana://<id>
     - jira://<id>
     - parent://<id>

## Examples

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
