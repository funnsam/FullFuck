@echo off
set commit=%1
git add .
git commit -m %commit%
git push origin main-pending
@echo on