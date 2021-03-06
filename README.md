# DoTask

my first go program to do some task journaling

## Commands

    dotask (l)ist|(s)how
    dotask (n)ow <title>
    dotask (c)clone|continue <ID> [now|<timestamp>]
    dotask delete <ID>
    dotask shutdown [now|<timestamp>]

    dotask <ID> asis|now|<timestamp> [<title> ...]

    <title> parses for:
     - asana://<id>
     - jira://<id>
     - parent://<id>

## Examples

    dotask now admin work asana://123456789
    20141010-001 | 2014-10-10, 10:07 | admin work | asana://123456789

    dotask 001 10:05 Admin: daily dtuff ...
    20141010-001 | 2014-10-10, 10:05 | Admin: daily dtuff ... | asana://123456789

    dotask 001 asis Admin: daily stuff ...
    20141010-001 | 2014-10-10, 10:05 | Admin: daily stuff ... | asana://123456789

    dotask now asana://123456790
    20141010-002 | 2014-10-10, 10:15 | title of the asna task | asana://123456790

    dotask continue 001
    20141010-003 | 2014-10-10, 12:45 | Admin: daily stuff ... | asana://123456789

    dotask 0 now lunch
    20141010-004 | 2014-10-10, 13:05 | lunch

    dotask continue 001
    20141010-005 | 2014-10-10, 13:35 | Admin: daily stuff ... | asana://123456789

    dotask continue 002
    20141010-006 | 2014-10-10, 14:00 | title of the asna task | asana://123456790

    dotask 0 17:00 (private)
    20141010-007 | 2014-10-10, 17:00 | (private)

    dotask continue 006
    20141010-008 | 2014-10-10, 17:33 | title of the asna task | asana://123456790

    dotask 008 17:35
    20141010-008 | 2014-10-10, 17:35 | title of the asna task | asana://123456790

    dotask shutdown
    20141010-009 | 2014-10-10, 20:00 | shutdown

    dotask list
    20141010-001 | 2014-10-10, 10:05 | Admin: daily stuff ... | asana://123456789
    20141010-002 | 2014-10-10, 10:15 | title of the asna task | asana://123456790
    20141010-003 | 2014-10-10, 12:45 | Admin: daily stuff ... | asana://123456789
    20141010-004 | 2014-10-10, 13:05 | lunch
    20141010-005 | 2014-10-10, 13:35 | Admin: daily stuff ... | asana://123456789
    20141010-006 | 2014-10-10, 14:00 | title of the asna task | asana://123456790
    20141010-007 | 2014-10-10, 17:00 | (private)
    20141010-008 | 2014-10-10, 17:35 | title of the asna task | asana://123456790
    20141010-009 | 2014-10-10, 20:45 | shutdown

... examples will look somehow different with ongoing implementation. ;-)
