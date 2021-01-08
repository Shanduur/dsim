clear all; close all; clc;

PLOTNAME = 'plot-amdahl.png';
TITLE = 'Results comparison';

threads = [ 1 2 4 8 16 ];

set_size = 80;

time = 4*60*60+13*60+32;
time_per_item = time/set_size;

for i = 1:length(threads)
  proportion(i) = (set_size - 1)/set_size;
  speedup(i)= 1 / (( 1 - proportion(i)) + (proportion(i)/threads(i)));
endfor

##actual_speedup = [ 1.0000 1.2836 1.8908 ];
actual_speedup = [ 1.0000 1.9565 2.8398 3.4613 3.9957 ];

error = actual_speedup - speedup;

disp(error)

fig = figure();

plot(threads, speedup, '-r+');
hold on;
plot(threads, actual_speedup, '-b+');
xlabel('Threads');
ylabel('Speedup');
title(TITLE);
legend('Calculated speedup', 'Actual speedup', 'location', 'northwest');
grid on;
print(fig, PLOTNAME);