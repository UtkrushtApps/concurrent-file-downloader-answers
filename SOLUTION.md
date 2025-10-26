# Solution Steps

1. Define a struct (DownloadResult) to carry download completion info, including file name and status.

2. Implement the simulateDownload function: it takes a file name, a channel for DownloadResult, and a WaitGroup pointer.

3. Inside simulateDownload, defer wg.Done(), sleep for 2-3 random seconds, and then send a DownloadResult with file name and status='completed' on the channel.

4. In main, seed the random number generator.

5. Create a list of four filenames to 'download'.

6. Create an unbuffered channel for DownloadResult and a sync.WaitGroup variable.

7. For each file, increment the WaitGroup and launch a goroutine calling simulateDownload.

8. Start a dedicated goroutine in main to close the completion channel once wg.Wait() ends (all downloads complete).

9. In main, use a for ... range loop over the channel to print download completion messages as they arrive.

10. After receiving all messages and the channel is closed, exit cleanly.

