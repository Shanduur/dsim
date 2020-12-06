param (
    [Parameter(Mandatory=$true)][string]$folder
)

$options = "-o=.\out\$(Get-Date -Format "yyyyMMddHHmmss") -uname=user -pwd=password -log-level=5"

$items = Get-ChildItem -name -Path $folder

New-Item -ItemType Directory -Force -Path .\out\
Get-ChildItem -Path .\out\ -Include *.* -File -Recurse | foreach { $_.Delete()}

echo "" > time.txt

$items | Sort-Object {Get-Random} | foreach {
    $query = $_
    $items | Sort-Object {Get-Random} | foreach {
        $train = $_

        $time = Measure-Command -Expression {.\client.exe $query $train}

        echo $time >> time.txt
    }
}
