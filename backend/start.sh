
#!/bin/bash
# Start the Docker daemon in the backgroun
# sleep 20
# cd /app/ch1-git-katas
# # Docker commands
# docker kill validator;
# docker kill gitkatas;
# docker rm validator;
# docker rm gitkatas;
# docker run --name="validator" -v $(pwd)/exercise:/exercise -d daafonsecato/validator:v1; 
# docker run --name="gitkatas" -p 7681:7681 -p 9090:9090 -v $(pwd)/exercise:/home/git-katas-user/exercise -d daafonsecato/gitkatas:v1; 
# docker exec -u root gitkatas bash -c 'chmod o+rx /var/hidden; chown -R git-katas-user:git-katas-user /home/git-katas-user/exercise;  cd /home/git-katas-user; su - git-katas-user -c \". /var/hidden/git-katas/basic-cleaning/setup.sh\";'; 
# docker exec -u root validator bash -c 'chown -R git-katas-user:git-katas-user /exercise';

# Start your application
cd /app
air
