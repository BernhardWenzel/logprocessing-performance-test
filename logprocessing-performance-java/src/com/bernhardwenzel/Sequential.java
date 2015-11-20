package com.bernhardwenzel;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.Duration;
import java.time.Instant;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Sequential
{
    public static final Pattern pattern = Pattern.compile("^[^#]*#[15][15]1110+$");

    public static void main(String[] args)
    {
        Instant start = Instant.now();
        System.out.println("Started: " + start);
        try (InputStream in = Files.newInputStream(Paths.get("../logprocessing-performance-go/log-sample.txt"));
             BufferedReader reader = new BufferedReader(new InputStreamReader(in)))
        {
            String line;
            while ((line = reader.readLine()) != null)
            {
                final Matcher matcher = pattern.matcher(line);
                if (matcher.matches())
                {
                    System.out.print(String.format("Match: %s\n", line));
                }
            }
        }
        catch (IOException ex)
        {
            System.err.println(ex);
        }
        Instant end = Instant.now();
        System.out.println(Duration.between(start, end));
    }
}
