#!/usr/bin/perl

my $number_args = $#ARGV + 1;  
if ($number_args != 2) {  
   print "usage: split.pl path files_in_set\n";  
   exit;  
}

my $base_dir_qfn = $ARGV[0];
my $i = 0;
my $dir;
my $count = $ARGV[1];

opendir(my $dh, $base_dir_qfn)
   or die("Can'\''t open dir \"$base_dir_qfn\": $!\n");

while (defined( my $fn = readdir($dh) )) {
   next if $fn =~ /^(?:\.\.?|dir_\d+)\z/;

   my $qfn = "$base_dir_qfn/$fn";

   if ($i % $count == 0) {
      $dir_qfn = sprintf("%s/dir_%03d", $base_dir_qfn, int($i/$count)+1);
      mkdir($dir_qfn)
         or die("Can'\''t make directory \"$dir_qfn\": $!\n");
   }

   rename($qfn, "$dir_qfn/$fn")
      or do {
         warn("Can'\''t move \"$qfn\" into \"$dir_qfn\": $!\n");
         next;
      };

   ++$i;
}