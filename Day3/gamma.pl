#!/usr/bin/perl
use warnings;
# use strict;

my $filename = 'input';

open(FH, '<', $filename) or die $!;
my @digittracker;


while(<FH>){
    $line = $_;
    for( $y=0; $y<=length($line); $y++){
        if ( substr($line, $y,1) eq "1" ){
            $digitaltracker[$y][1] +=1;
        }else{
            $digitaltracker[$y][0] +=1;
        }
    }
}



for($x=0; $x < $#digitaltracker-1; $x++){
    if($digitaltracker[$x][0] gt $digitaltracker[$x][1]){
        $gamma .= "0";
        $epsilon .= "1";
    }else{
        $gamma .= "1";
        $epsilon .= "0";
    }
}



print "\ngamma rate: ", oct("0b".$gamma) , "\n";
print "epsilon rate: ", oct("0b".$epsilon) , "\n";

print "power consumption: ", oct("0b".$epsilon) * oct("0b".$gamma), "\n";
close(FH);