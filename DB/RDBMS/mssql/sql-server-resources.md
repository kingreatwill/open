# Technical Resources for SQL Server & more

Most open source, all free. Scripts, tools, and resources pertinent to SQL Server and beyond. 

## T-SQL Scripts
- [Ola's Maintenance Scripts](https://github.com/olahallengren/sql-server-maintenance-solution)
- [Brent's First Responder Kit](https://github.com/BrentOzarULTD/SQL-Server-First-Responder-Kit)
- [Glenn Berry's Diagnostic Queries](https://www.sqlskills.com/blogs/glenn/category/dmv-queries/)
- [Tiger Tool Box](https://github.com/Microsoft/tigertoolbox)
- [sp_WhoIsActive](http://whoisactive.com/downloads/)
- [sp_InEachDB](https://github.com/BrentOzarULTD/SQL-Server-First-Responder-Kit/blob/dev/sp_ineachdb.sql) - replacement for the undocumented [sp_MSForEachDB](http://sqlblog.com/blogs/aaron_bertrand/archive/2010/12/29/a-more-reliable-and-more-flexible-sp-msforeachdb.aspx)
- [Jeff Moden's DelimitedSplit8k](http://www.sqlservercentral.com/articles/Tally+Table/72993/) - string splitter function

## References

- [SQL Style Guide](http://www.sqlstyle.guide/)
- [T-SQL Style Guide](https://lowlydba.github.io/tsqlstyle.guide/)
- [DBATools Build Reference](https://sqlcollaborative.github.io/builds)
- [MSSQL Waitopedia](https://www.spotlightessentials.com/public/waitopedia)
- [Module Signing Info](https://modulesigning.info/)
- [SQL 2017 & Azure Database Permissions Poster](assets/Permissions_Poster_2017_and_SQLDB.PDF)
- [The Data Loading Performance Guide](https://docs.microsoft.com/en-us/previous-versions/sql/sql-server-2008/dd425070(v=sql.100))
- [Database Hiearachy of Monitoring Infographic](https://www.lowlydba.com/database-hierarchy-of-monitoring/)
- [Scheduling Powershell Tasks With SQL Agent](https://dbatools.io/agent/)

## SSMS Plugins
- [Apex SQL Refactor](http://www.apexsql.com/sql_tools_refactor.aspx)
- [Apex SQL Complete](http://www.apexsql.com/sql_tools_complete.aspx)
- [Apex SQL Search](http://www.apexsql.com/sql_tools_search.aspx)
- [SQL Sentry Plan Explorer](https://www.sentryone.com/plan-explorer/)
- [Dell Spotlight Essentials](https://www.spotlightessentials.com/spotlight-extensions)

## Tools
| Name | Type | Open Source | Author |
| ---- | ---- | ----------- | ------ |
| [DBA Tools](https://dbatools.io) | Admin | 👍 | [SQL Collaborative](https://dbatools.io/team/) |
| [DBA Checks](https://dbachecks.io) | Admin | 👍 | [SQL Collaborative](https://dbatools.io/team/) |
| [MinionWare](http://www.minionware.net/) | Admin | | [MinionWare](http://www.minionware.net/meet-the-team/)|
| [DLM Dashboard](http://www.red-gate.com/products/dlm/dlm-dashboard/) | Devops | | [Redgate](https://www.red-gate.com/) |
| [Flyway](https://flywaydb.org/) | Migrations | 👍 | [Axel Fontaine](https://axelfontaine.com/) |
| [Is It SQL?](http://www.scalesql.com/isitsql/) | Monitoring | | [Bill Graziano](http://www.scalesql.com/about.html)
| [Spotlight Cloud Basic](https://www.spotlightcloud.io/pricing) | Monitoring | | [Quest](https://www.quest.com/) |
| [SQL Watch](https://sqlwatch.io/) | Monitoring | 👍 | [Marcin Gminski](https://marcin.gminski.net/) |
| [SchemaZen](https://github.com/sethreno/schemazen#schemazen---script-and-create-sql-server-objects-quickly) | Scripting | :+1: | [Seth Reno](https://github.com/sethreno) |
| [mssql-scripter](https://github.com/Microsoft/sql-xplat-cli/) | Scripting | 👍 | Microsoft |
| [DBDiagram.io](https://dbdiagram.io/) | Sharing | | [holistics.io](https://www.holistics.io)|
| [Format Text as Table](https://senseful.github.io/text-table/) | Sharing | | [Senseful Solutions](https://senseful.github.io/) |
| [DBFit](http://www.methodsandtools.com/tools/dbfit.php) | Testing | | <a href="https://javornikolov.wordpress.com/">Yavor Nikolov</a> |
| [DB Fiddle](https://dbfiddle.uk/) | Testing/Sharing | | [Jack Douglas](https://douglastechnology.co.uk/) |
| [SQL Fiddle](http://sqlfiddle.com/) | Testing/Sharing | | [Jake Feasel](http://stackoverflow.com/users/808921/jake-feasel) |
| [Paste The Plan](https://pastetheplan.com/) | Tuning/Sharing | | [Brent Ozar Unlimited](https://www.brentozar.com/)

## Other
- [Wide World Importer Database](https://github.com/Microsoft/sql-server-samples) - Successor to the AdventureWorks sample database for SQL Server 2016+
- [Stack Overflow Database](https://www.brentozar.com/archive/2015/10/how-to-download-the-stack-overflow-database-via-bittorrent/) - Brent Ozar's packaging of the Stack Overflow database
- [idownvotedbecau.se](http://idownvotedbecau.se/) - Linkable downvoting rationale for Stack Exchange comments.
- [Open Source Pull Request Template](https://www.talater.com/open-source-templates/#/) - Whimsical choose-your-own-adventure that climaxes with a customized pull request template by [Tal Ater](https://twitter.com/TalAter)

## Good Reads

### AlwaysOn
- [The AlwaysOn Health Model Part 2 - Extending the Health Model](https://techcommunity.microsoft.com/t5/SQL-Server/The-AlwaysOn-Health-Model-Part-2-Extending-the-Health-Model/ba-p/384043?advanced=false&collapse_discussion=true&q=the%20alwayson%20health%20model&search_type=thread)
- [When Always On Isn't: Handling Outages In Your Application](https://www.brentozar.com/archive/2017/01/always-isnt-handling-outages-application/) by Brent Ozar
- [SQL Server 2016/2017: Availability group secondary replica redo model and performance](https://blogs.msdn.microsoft.com/sql_server_team/sql-server-20162017-availability-group-secondary-replica-redo-model-and-performance/) - Detailed information on parallelism in redo worker threads

### Misc.
- [Modern Data Analysis: Don't Trust Your Spreadsheet][betterment]
- [T-SQL Interview Questions](https://www.mssqltips.com/sqlservertip/1450/sql-server-developer-tsql-interview-questions/) by [Jeremy Kadlec](https://www.mssqltips.com/sqlserverauthor/38/jeremy-kadlec/)
- [Developer Interview Questions](https://www.brentozar.com/archive/2009/06/top-10-developer-interview-questions-about-sql-server/) by Brent Ozar
- [Tuning Cost Threshold](http://sqlblog.com/blogs/jonathan_kehayias/archive/2010/01/19/tuning-cost-threshold-of-parallelism-from-the-plan-cache.aspx)
- [5 Rules of Normalization][normrules] by Marc Rettig
- [T-SQL Code Smells][smelly] by [Phil Factor][phil]
- [Fighting Evil in Your Code: Comments on Comments](https://www.red-gate.com/simple-talk/opinion/opinion-pieces/fighting-evil-code-comments-comments/) by [Michael Sorens](https://www.red-gate.com/simple-talk/author/michael-sorens/)
- [The Security of Modern Password Expiration: An Algorithmic Framework and Empirical Analysis](https://www.cs.unc.edu/~reiter/papers/2010/CCS.pdf)

[betterment]: https://www.betterment.com/resources/inside-betterment/engineering/modern-data-analysis-dont-trust-your-spreadsheet/
  "Betterment Blog"
[isitsql]: http://www.scalesql.com/isitsql/
  "Is It SQL?"
[schemazen]: https://github.com/sethreno/schemazen#schemazen---script-and-create-sql-server-objects-quickly
  "SchemaZen"
[dbfit]: http://www.methodsandtools.com/tools/dbfit.php
  "DB Fit"
[fiddle]: http://sqlfiddle.com/
  "SQL Fiddle"
[normrules]: https://github.com/LowlyDBA/miscmssql/blob/master/Best%20Practices/Marc_Rettig_5_Rules_of_Normalization_Poster.pdf
  "5 Rules of Normalization"
 [smelly]: https://www.red-gate.com/simple-talk/sql/t-sql-programming/sql-code-smells/
 [phil]: https://www.red-gate.com/simple-talk/author/phil-factor/
