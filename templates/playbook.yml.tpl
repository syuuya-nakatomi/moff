---
- hosts: all
  become: yes
  gather_facts: no

  tasks:
  {{ range . }}
  - name: Deploy {{ . }}
    yum:
      name: "{{ . }}"
      state: latest
    tags:
      - deploy
  {{ end }}