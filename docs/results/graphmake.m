clear all; close all; clc;

UPPERBOUND = 31;
FILE = 'results-80.csv';
PLOTNAME1 = 'plot-80.png';
TITLE1 = 'Processing time and thread count relationship - 80 small files';
PLOTNAME2 = 'plot-speedup-80.png';
TITLE2 = 'Speedup and thread count relationship - 80 small files';
##FILE = 'results-4.csv';
##PLOTNAME1 = 'plot-4.png';
##TITLE1 = 'Processing time and thread count relationship - 4 big files';
##PLOTNAME2 = 'plot-speedup-4.png';
##TITLE2 = 'Speedup and thread count relationship - 4 big files';
##FILE = 'results-4n.csv';
##PLOTNAME1 = 'plot-4n.png';
##TITLE1 = 'Processing time and thread count relationship - 4 big files';
##PLOTNAME2 = 'plot-speedup-4n.png';
##TITLE2 = 'Speedup and thread count relationship - 4 big files';
##UPPERBOUND = 7;

THREAD_COL = uint8(1);
ELAPSED_HMS_COL = uint8(2);
IDEAL_HMS_COL = uint8(3);
ELAPSED_S_COL = uint8(4);
IDEAL_S_COL = uint8(5);
SYSTEM_COL = uint8(6);
USER_COL = uint8(7);
CPU_COL = uint8(8);

pkg load io;


  
data = csv2cell(FILE);

threads = cell2mat(data([1:UPPERBOUND], THREAD_COL));
[~, unique_threads_id] = unique(threads,'first');
elapsed_times = cell2mat(data([1:UPPERBOUND], ELAPSED_S_COL));
linear_times = cell2mat(data([1:UPPERBOUND], IDEAL_S_COL));

tmp = 1;
id1 = 0;
id2 = 1;
for i = 1:UPPERBOUND
    if threads(i) == tmp
      id1++;
      ids(id1) = i;
    else
      tmp = threads(i);
      elapsed_avg(id2) = mean(elapsed_times(ids));
      elapsed_max(id2) = abs(elapsed_avg(id2) - max(elapsed_times(ids)));
      elapsed_min(id2) = abs(elapsed_avg(id2) - min(elapsed_times(ids)));
      e_max(id2) = max(elapsed_times(ids));
      e_min(id2) = min(elapsed_times(ids));
      id2++;
      ids = 0;
      id1 = 1;
      ids(id1) = i;
    endif
endfor;
elapsed_avg(id2) = mean(elapsed_times(ids));
elapsed_max(id2) = abs(elapsed_avg(id2) - max(elapsed_times(ids)));
elapsed_min(id2) = abs(elapsed_avg(id2) - min(elapsed_times(ids)));
e_max(id2) = max(elapsed_times(ids));
e_min(id2) = min(elapsed_times(ids));

for i = 1:length(elapsed_avg)
  speedup(i) = elapsed_avg(1) / elapsed_avg(i);
  speedup_min(i) = abs(speedup(i) - (elapsed_avg(1) / e_min(i)));
  speedup_max(i) = abs(speedup(i) - (elapsed_avg(1) / e_max(i))); 
endfor

for i = 1:UPPERBOUND
  linear_speedup(i) = linear_times(1)/linear_times(i);
endfor

fig = figure();

plot(threads(unique_threads_id), linear_times(unique_threads_id), '-r+');
hold on;
errorbar(threads(unique_threads_id), elapsed_avg, elapsed_max, elapsed_min, '~-b');
xlabel('Threads');
ylabel('Processing time [s]');
title(TITLE1);
legend('Linear times', 'Elapsed times', 'location', 'northwest');
grid on;
print(fig, PLOTNAME1);
hold off;

plot(threads(unique_threads_id), linear_speedup(unique_threads_id), '-r+');
hold on;
errorbar(threads(unique_threads_id), speedup, speedup_max, speedup_min, '~-b');
xlabel('Threads');
ylabel('Speedup');
title(TITLE2);
legend('Linear speedup', 'Actual speedup', 'location', 'northwest');
grid on;
print(fig, PLOTNAME2);
