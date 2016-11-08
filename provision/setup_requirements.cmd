echo 'Ensuring Chocolatey is Installed'
@powershell -NoProfile -ExecutionPolicy Bypass -File "c:\vagrant_data\provision\install_choco.ps1"
choco install osquery -y
