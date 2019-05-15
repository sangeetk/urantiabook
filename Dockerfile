FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD urantiabook urantiabook
EXPOSE 8080
CMD ["/urantiabook"]
