package com.bernhardwenzel;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.Duration;
import java.time.Instant;

class BufferedReaderPerformance
{
    public static void main(String[] args)
    {
        Instant start = Instant.now();
        System.out.println("Started: " + start);
        try (InputStream in = Files.newInputStream(Paths.get("../logprocessing-performance-go/log-sample.txt"));
             BufferedReader reader = new BufferedReader(new InputStreamReader(in)))
        {
            String line;
            int i = 0;
            while ((line = reader.readLine()) != null)
            {
                i++;
            }
            System.out.println(i);
        }
        catch (IOException ex)
        {
            System.err.println(ex);
        }
        Instant end = Instant.now();
        System.out.println(Duration.between(start, end));
    }
}
