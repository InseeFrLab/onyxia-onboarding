FROM alpine  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ADD onyxia-onboarding /root/onyxia-onboarding  
CMD [ "/bin/sh", "-c", "./onyxia-onboarding" ]