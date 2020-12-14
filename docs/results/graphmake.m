clear all; close all; clc;

#PLOTNAME = "plot-80.png";
#FILE = "results-80.csv";
#TITLE = "Time and thread count relationship - 80 small files.";
#UPPERBOUND = 31;
PLOTNAME = "plot-4.png";
FILE = "results-4.csv";
TITLE = "Time and thread count relationship - 4 big files.";
UPPERBOUND = 7;

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
ideal_times = cell2mat(data([1:UPPERBOUND], IDEAL_S_COL));

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
      id2++;
      ids = 0;
      id1 = 1;
      ids(id1) = i;
    endif
endfor;
elapsed_avg(id2) = mean(elapsed_times(ids));
elapsed_max(id2) = abs(elapsed_avg(id2) - max(elapsed_times(ids)));
elapsed_min(id2) = abs(elapsed_avg(id2) - min(elapsed_times(ids)));

fig = figure();
scatter(threads, ideal_times);
plot(threads(unique_threads_id), ideal_times(unique_threads_id));
hold on;
errorbar(threads(unique_threads_id), elapsed_avg, elapsed_max, elapsed_min, '~');
xlabel("Threads");
ylabel("Processing time [s]");
title(TITLE);
legend("Ideal times", "Elapsed times");
print(fig, PLOTNAME);

