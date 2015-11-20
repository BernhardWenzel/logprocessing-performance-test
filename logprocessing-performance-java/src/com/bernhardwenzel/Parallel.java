package com.bernhardwenzel;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.Duration;
import java.time.Instant;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Parallel
{
    public static final Pattern pattern = Pattern.compile("^[^#]*#[15][15]1110+$");
    static final int numberOfWorkers = 2;
    static ExecutorService executor = Executors.newFixedThreadPool(numberOfWorkers);
    static final ArrayBlockingQueue<String> queue = new ArrayBlockingQueue<String>(512);

    public static void main(String[] args)
    {
        Instant start = Instant.now();
        System.out.println("Started: " + start);
        try (
                InputStream in = Files.newInputStream(Paths.get("../logprocessing-performance-go/log-sample.txt"));
                BufferedReader reader = new BufferedReader(new InputStreamReader(in))
        )
        {
            // start workers
            for (int i = 0; i < numberOfWorkers; i++)
            {
                Worker worker = new Worker(queue, i);
                executor.submit(worker);
            }

            String line;
            while ((line = reader.readLine()) != null)
            {
                queue.put(line);
            }
            executor.shutdownNow();
        }
        catch (Exception ex)
        {
            System.err.println(ex);
        }
        System.out.println("Done");
        Instant end = Instant.now();
        System.out.println(Duration.between(start, end));
    }

    static class Worker implements Runnable
    {
        private final BlockingQueue<String> queue;
        private final int id;

        public Worker(BlockingQueue<String> q, int id)
        {
            queue = q;
            this.id = id;
        }

        public void run()
        {
            try
            {
                while (!Thread.currentThread().isInterrupted())
                {
                    String line = queue.take();
                    final Matcher matcher = pattern.matcher(line);
                    if (matcher.matches())
                    {
                        System.out.print(String.format("[%s] Match: %s\n", id, line));
                    }
                }
            }
            catch (InterruptedException ex)
            {
            }
        }
    }
}
