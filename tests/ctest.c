#include <stdio.h>
#include "revisions.h"
int main(int argc, char **argv)
{
    printf("Version %d.%d-%d %s\n" , VERSION_MAJOR, VERSION_MINOR, VERSION_BUILD  , BUILD_TIME );
    printf("REPO %s\n" , REPO_URL);
    printf("BRANCH %s\n",BRANCH_NAME);
    printf("LONG COMMIT ID %s\n" , LONG_COMMIT_ID);
}