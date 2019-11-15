# srctrace - generate source code details for traceability

## Background

Applications distributed in binary form should be traceable back to its source. When using a revision control system like git, this would mean being able to extract the source regardless of how old or home any revisions it has undergone since.

This projectlet generates a fragment of source code that can be compiled into the application and retrieved at runtime to report the details as necessary.

## Usage

        usage: srctrace [-h|--help] [-v|--verbose] [-m|--major <integer>] [-n|--minor
                        <integer>] [-b|--build <integer>] [-L|--language
                        (C|C++|Ada|python|go)] [-o|--output "<value>"]

                        generate source trace info

        Arguments:

        -h  --help      Print help information
        -v  --verbose   Verbose. Default: false
        -m  --major     Major version. Default: 0
        -n  --minor     Minor version. Default: 0
        -b  --build     Build Number. Default: 999
        -L  --language  Language to output
        -o  --output    Output file base name. Default: revisions

## Examples

### C Application

        ../../bin/srctrace --language C -m 2 -n 3 -b 234
        2019/11/07 05:46:03 srctrace
        $ cat revisions.h
        // C header generator
        // File: revisions.h
        #define BUILD_TIME "Thu Nov 7 2019 05:46:03"
        #define VERSION_MAJOR (2)
        #define VERSION_MINOR (3)
        #define VERSION_BUILD (234)
        #define REPO_URL "git@gitlab.com:privatetutor/projectlets/go.git"
        #define BRANCH_NAME "master"
        #define SHORT_COMMIT_ID "cefa722"
        #define LONG_COMMIT_ID “cefa72267cc4d3a07fbf5e672b0053116d582aa7”

### Ada Application

        ../../bin/srctrace --language Ada -m 2 -n 3 -b 234
        2019/11/07 05:48:39 srctrace
        $ cat revisions.ads
        package revisions is
        -- Ada spec generator
        -- File: revisions.ads
            BUILD_TIME : constant String := "Thu Nov 7 2019 05:48:39" ;
            VERSION_MAJOR : constant := 2 ;
            VERSION_MINOR : constant := 3 ;
            VERSION_BUILD : constant := 234 ;
            REPO_URL : constant String := "git@gitlab.com:privatetutor/projectlets/go.git" ;
            BRANCH_NAME : constant String := "master" ;
            SHORT_COMMIT_ID : constant String := "cefa722" ;
            LONG_COMMIT_ID : constant String := "cefa72267cc4d3a07fbf5e672b0053116d582aa7" ;
        end revisions ;
