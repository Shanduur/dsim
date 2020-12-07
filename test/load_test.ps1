param (
    [Parameter(Mandatory=$true)][string]$folder
)

$options = "-o=.\out\$(Get-Date -Format "yyyyMMddHHmmss") -uname=user -pwd=password -log-level=5"

$items = Get-ChildItem -name -Path $folder

New-Item -ItemType Directory -Force -Path .\out\
New-Item -ItemType Directory -Force -Path .\log\
Get-ChildItem -Path .\out\ -Include *.* -File -Recurse | foreach { $_.Delete()}

echo "" > time.txt

$items | Sort-Object {Get-Random} | foreach {
    $s1 = $_
    $items | Sort-Object {Get-Random} | foreach {
        $s2 = $_

        $time = Measure-Command -Expression {.\client.exe "-source-img1=$s1 -source-img2=$s2 $options"}

        echo $time >> .\log\time.txt
    }
}
