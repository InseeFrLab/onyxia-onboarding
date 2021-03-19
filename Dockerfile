FROM alpine  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ADD onyxia-onboarding /root/onyxia-onboarding  
RUN chmod +x /root/onyxia-onboarding  
CMD [ "./onyxia-onboarding" ]