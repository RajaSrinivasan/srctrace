with Text_Io; use Text_Io;
with Revisions; use Revisions;

procedure Test is
begin
   Put_Line(BUILD_TIME);
   Put_Line(REPO_URL);
   Put_Line("Branch " + BRANCH_NAME );
   Put_Line("Long Commit Id " + LONG_COMMIT_ID );
end Test ;
