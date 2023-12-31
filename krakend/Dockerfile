FROM devopsfaith/krakend:latest

# Copy the KrakenD configuration file to the container
COPY krakend.json /etc/krakend/krakend.json

# Expose the KrakenD port
EXPOSE 8080

# Command to run KrakenD
CMD ["krakend", "run", "-c", "/etc/krakend/krakend.json"]
