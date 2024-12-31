FROM confluentinc/cp-kafka:7.3.0

COPY --chmod=0755 create-topics.sh /create-topics.sh

CMD ["/create-topics.sh"]
