#!/bin/sh

if [ "x$ENV" == "x" ]; then
   ENV="dev"
fi
echo "Deploying in $ENV"
/opt/cb-quiz/cb-quiz --env $ENV > /opt/cb-quiz/log 2>&1 
tail -f /dev/null
