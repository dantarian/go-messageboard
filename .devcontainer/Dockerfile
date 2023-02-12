FROM qmcgaw/godevcontainer

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=${USER_UID}
ARG EXTHOME

RUN addgroup -g ${USER_GID} ${USERNAME} \
  && adduser --uid ${USER_UID} --ingroup ${USERNAME} -D ${USERNAME}
RUN apk add shadow \
  && usermod --shell /bin/zsh ${USERNAME} \
  && apk del shadow
RUN cp -r /root/. /home/${USERNAME} \
  && chown -R ${USER_UID}:${USER_GID} /home/${USERNAME} \
  && chown -R ${USER_UID}:${USER_GID} /go \
  && chmod a+rx /usr/local/bin/*
RUN sed -i -e s/root/home\\/${USERNAME}/ /home/${USERNAME}/.zshrc
RUN apk add --no-cache make

USER ${USERNAME}